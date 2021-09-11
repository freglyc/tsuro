package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
)

// Websocket upgrade
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	hub  *Hub            // Hub
	conn *websocket.Conn // Websocket connection
	send chan []byte     // Buffered channel of outbound messages
}

// Pumps messages from the websocket connection to the hub
func (client *Client) readPump(wg *sync.WaitGroup) {
	defer func() {
		client.hub.unregister <- client
		_ = client.conn.Close()
	}()
	client.conn.SetReadLimit(maxMessageSize)
	_ = client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { _ = client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	wg.Done()
	for {
		var msg Message
		err := client.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("game read errored with error %v\n", err)
			}
			break
		}
		// Error check msg
		if (msg.Kind == "place" && (msg.Row < 0 || msg.Col < 0 || msg.Row > msg.Size-1 || msg.Col > msg.Size-1 ||
			msg.Idx < 0 || msg.Idx > 2)) || ((msg.Kind == "rotateRight" || msg.Kind == "rotateLeft") &&
			(msg.Idx < 0 || msg.Idx > 2)) || msg.Players < 2 || msg.Players > 8 || msg.Size != 6 || msg.GameID == "" {
			log.Println("INVALID MESSAGE")
		} else {
			client.hub.broadcast <- ClientMessage{
				Client:  client,
				Message: msg,
			}
		}
	}
}

// Pumps messages from the hub to the websocket connection
func (client *Client) writePump(wg *sync.WaitGroup) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = client.conn.Close()
	}()
	wg.Done()
	for {
		select {
		case message, ok := <-client.send:
			_ = client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			_ = client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func ServeWs(hub *Hub, response http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go client.readPump(wg)
	go client.writePump(wg)
	wg.Wait()
	hub.register <- client
}
