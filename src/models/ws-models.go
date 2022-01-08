package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

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
