package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Websocket upgrade
var upgrader = websocket.Upgrader{
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
func (client *Client) readPump() {
	defer func() {
		client.hub.unregister <- client
		_ = client.conn.Close()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	_ = client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { _ = client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var msg Message
		err := client.conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			break
		}
		client.hub.broadcast <- ClientMessage{
			Client:  client,
			Message: msg,
		}
	}
}

// Pumps messages from the hub to the websocket connection
func (client *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				_ = client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
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
	hub.register <- client

	go client.readPump()
	go client.writePump()
}
