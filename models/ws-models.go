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
	Key string `json:"data"`
}

type UserJoinsMessage struct {
	ID     int    `json:"id"`
	Sprite string `json:"sprite"`
}

type BroadcastMessage struct {
	Type        WSMessageType `json:"type"`
	Players     []Player      `json:"players"`
	Enemies     []Player      `json:"enemies"`
	Drops       []Drop        `json:"drops"`
	Projectiles []Projectile  `json:"projectiles"`
}

type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Player *Player
}

type Hub struct {
	Clients           map[*Client]bool
	Register          chan *Client
	Unregister        chan *Client
	Broadcast         chan bool
	PlayerSprites     []Sprite
	EnemySprites      []Sprite
	Enemies           []*Player
	DropSprites       []DropSprite
	Drops             []*Drop
	ProjectileSprites []ProjectileSprite
	Projectiles       map[*Projectile]bool
	FloorLayer        Layer
	MapArea           Area
}

func (h *Hub) GetAliveEnemies(excludeId int) []*Player {
	var enemies []*Player
	for _, enemy := range h.Enemies {
		if enemy.ID == excludeId || enemy.Dead {
			continue
		}
		enemies = append(enemies, enemy)
	}
	return enemies
}

func (h *Hub) GetAlivePlayers(excludeId int) []*Player {
	var players []*Player
	for client := range h.Clients {
		players = append(players, client.Player)
	}
	for _, player := range players {
		if player.ID == excludeId || player.Dead {
			continue
		}
		players = append(players, player)
	}
	return players
}
