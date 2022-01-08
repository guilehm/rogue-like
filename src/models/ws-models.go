package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type WSMessageType string

type WSMessage struct {
	MessageType WSMessageType   `json:"type"`
	Data        json.RawMessage `json:"data"`
}

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Player *Player
}

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
}

func (hub *Hub) Start() {
	fmt.Println("starting")
	for {
		select {
		case client := <-hub.Register:
			log.Println("registering client")
			hub.Clients[client] = true
		case client := <-hub.Unregister:
			if _, ok := hub.Clients[client]; ok {
				delete(hub.Clients, client)
			}
		}
	}
}
