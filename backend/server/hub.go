package server

import (
	"encoding/json"
	"github.com/freglyc/tsuro/game"
	"log"
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
}

func NewHub() *Hub {
	return &Hub{
		games:      make(map[string]*GameHandler),
		clients:    make(map[*Client]string),
		broadcast:  make(chan ClientMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			log.Println("CLIENT REGISTERED")
			hub.clients[client] = ""
		case client := <-hub.unregister:
			log.Println("CLIENT UNREGISTERED")
			gameID := hub.clients[client]
			if gameID != "" {
				if handler := hub.games[gameID]; handler != nil {
					delete(handler.Clients, client)
					if len(handler.Clients) == 0 {
						delete(hub.games, gameID)
					}
				}
			}
			delete(hub.clients, client)
			close(client.send)
		case clientMessage := <-hub.broadcast:
			handler := hub.games[clientMessage.GameID]
			if handler == nil {
				handler = NewHandler(tsuro.NewGame(clientMessage.GameID, clientMessage.Options))
			}
			switch clientMessage.Kind {
			case "join":
				// Leave other game if in one
				oldID := hub.clients[clientMessage.Client]
				if oldID != "" {
					if handler := hub.games[oldID]; handler != nil {
						delete(handler.Clients, clientMessage.Client)
						if len(handler.Clients) == 0 {
							delete(hub.games, oldID)
						}
					}
				}
				// Add to game
				hub.clients[clientMessage.Client] = clientMessage.GameID
				handler.Clients[clientMessage.Client] = true
			case "rotateRight":
				handler.Game.RotateRight(tsuro.Team(clientMessage.Team), clientMessage.Idx)
			case "rotateLeft":
				handler.Game.RotateLeft(tsuro.Team(clientMessage.Team), clientMessage.Idx)
			case "place":
				handler.Game.Place([]int{clientMessage.Row, clientMessage.Col}, tsuro.Team(clientMessage.Team), clientMessage.Idx)
			case "reset":
				handler.Game.Reset()
			default:
				log.Print("Not a valid message")
			}

			data, err := json.Marshal(handler)
			if err != nil {
				log.Println(err)
				return
			}

			for client := range handler.Clients {
				select {
				case client.send <- data:
				default:
					close(client.send)
					delete(handler.Clients, client)
					if len(handler.Clients) == 0 {
						delete(hub.games, clientMessage.GameID)
					}
				}
			}
		}
	}
}
