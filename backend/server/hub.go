package server

import (
	"encoding/json"
	"github.com/freglyc/tsuro/game"
	"log"
	"time"
)

type Message struct {
	GameID string     `json:"gameID"`
	Kind   string     `json:"kind"`
	Team   tsuro.Team `json:"team"`
	Idx    int        `json:"idx"`
	Row    int        `json:"row"`
	Col    int        `json:"col"`
	tsuro.Options
}

type ClientMessage struct {
	Client *Client
	Message
}

type Hub struct {
	games      map[string]*GameHandler // mapping of game id to handler
	clients    map[*Client]string      // map of connected clients to gameID, gameID is "" if not joined
	broadcast  chan ClientMessage      // inbound messages from client connections
	register   chan *Client            // register requests from clients to the hub
	unregister chan *Client            // unregister requests from clients to the hub
	remove     chan string             // removes a game
}

func NewHub() *Hub {
	return &Hub{
		games:      make(map[string]*GameHandler),
		clients:    make(map[*Client]string),
		broadcast:  make(chan ClientMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		remove:     make(chan string),
	}
}

func (hub *Hub) Run() {
	go hub.clean()
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = ""
		case client := <-hub.unregister:
			if !byteChanIsClosed(client.send) {
				close(client.send)
			}
			delete(hub.clients, client)
		case gameID := <-hub.remove:
			delete(hub.games, gameID)
		case clientMessage := <-hub.broadcast:
			handler := hub.games[clientMessage.GameID]
			if handler == nil {
				handler = NewHandler(tsuro.NewGame(clientMessage.GameID, clientMessage.Options), hub)
				hub.games[clientMessage.GameID] = handler
			}
			switch clientMessage.Kind {
			case "join":
				hub.clients[clientMessage.Client] = clientMessage.GameID
			case "rotateRight":
				handler.Game.RotateRight(clientMessage.Team, clientMessage.Idx)
			case "rotateLeft":
				handler.Game.RotateLeft(clientMessage.Team, clientMessage.Idx)
			case "place":
				handler.Game.Place([]int{clientMessage.Row, clientMessage.Col}, clientMessage.Team, clientMessage.Idx)
				if handler.Game.Time > 0 && handler.Game.Started {
					if len(handler.Game.Winner) > 0 {
						handler.StopTimer()
					} else {
						handler.StartTimer()
					}
				}
			case "reset":
				handler.StopTimer()
				handler.Game.Reset()
			default:
				log.Print("Not a valid message")
			}
			handler.UpdateTime() // Update countdown clock if there is one
			data, err := json.Marshal(handler)
			if err != nil {
				log.Println(err)
				return
			}
			for client, gameID := range hub.clients {
				if gameID == clientMessage.GameID {
					select {
					case client.send <- data:
					default:
						if !byteChanIsClosed(client.send) {
							close(client.send)
						}
						delete(hub.clients, client)
					}
				}
			}
		}
	}
}

func (hub *Hub) clean() {
	// every hour remove 3 hour old games
	for range time.Tick(time.Hour) {
		for gameID, gameServer := range hub.games {
			deleteTime := gameServer.CreatedAt.Add(time.Duration(3) * time.Hour)
			if time.Now().After(deleteTime) {
				hub.remove <- gameID
			}
		}
	}
}

func byteChanIsClosed(ch <-chan []byte) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
