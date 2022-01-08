package models

import "github.com/gorilla/websocket"

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
}

type Hub struct {
	Clients map[*Client]bool
}
