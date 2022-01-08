package models

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type WSMessageType string

type WSMessage struct {
	MessageType WSMessageType   `json:"type"`
	Data        json.RawMessage `json:"data"`
}

var (
	UserJoins WSMessageType = "user-joins"
	Broadcast WSMessageType = "broadcast"
	KeyDown   WSMessageType = "key-down"
)

type KeyDownMessage struct {
	Key KeyName `json:"key"`
}

type UserJoinsMessage struct {
	Sprite string `json:"sprite"`
}

type BroadcastMessage struct {
	Type    WSMessageType `json:"type"`
	Players []*Player     `json:"players"`
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
	Broadcast  chan bool
	Sprites    []Sprite
}
