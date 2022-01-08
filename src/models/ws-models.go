package models

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gorilla/websocket"
)

type WSMessageType string

type WSMessage struct {
	MessageType WSMessageType   `json:"type"`
	Data        json.RawMessage `json:"data"`
}

var (
	UserJoins WSMessageType = "user-joins"
)

type UserJoinsMessage struct {
	Sprite string `json:"sprite"`
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
	Sprites    []Sprite
}

func (hub *Hub) GetSprite(name SpriteName) (Sprite, error) {
	for _, sprite := range hub.Sprites {
		if sprite.Name == name {
			return sprite, nil
		}
	}
	return Sprite{}, errors.New("sprite not found")
}

func (hub *Hub) createSprites() {
	hub.Sprites = []Sprite{
		{
			Name:         Warrior,
			TileSet:      Characters,
			SpriteX:      0,
			SpriteY:      0,
			SpriteWidth:  8,
			SpriteHeight: 8,
			HP:           100,
			MoveRange:    1,
			AttackRange:  1,
		},
	}
}

func (hub *Hub) Start() {
	hub.createSprites()
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