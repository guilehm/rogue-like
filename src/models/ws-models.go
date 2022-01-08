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
	Clients map[*Client]bool
}

func (hub *Hub) Start() {
	fmt.Println("starting")
}
