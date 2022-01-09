package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"rogue-like/models"
	"rogue-like/services"
	"time"

	"github.com/gorilla/websocket"
)

func handleKeyDown(client *models.Client, message models.WSMessage) error {
	key := ""
	err := json.Unmarshal(message.Data, &key)
	if err != nil {
		log.Println("error during unmarshall:", err)
		return err
	}

	client.Player.Move(key)
	// TODO: only set last interaction if player actually moved
	for c := range client.Hub.Clients {
		c.Player.LastInteraction = false
	}
	client.Player.LastInteraction = true
	client.Hub.Broadcast <- true

	return nil
}

func handleUserJoins(
	s *services.GameService, conn *websocket.Conn, client *models.Client, quit chan bool, message models.WSMessage,
) error {

	data := models.UserJoinsMessage{}
	err := json.Unmarshal(message.Data, &data)
	if err != nil {
		log.Println("error during unmarshall:", err)
		return err
	}

	rand.Seed(time.Now().UnixNano())
	sprite := s.Hub.PlayerSprites[rand.Int()%len(s.Hub.PlayerSprites)]
	enemySprite := s.Hub.EnemySprites[rand.Int()%len(s.Hub.EnemySprites)]

	s.Hub.Enemies = append(s.Hub.Enemies, &models.Enemy{
		Sprite:          enemySprite,
		Health:          enemySprite.HP,
		PositionX:       8 * 7,
		PositionY:       8 * 9,
		LastPosition:    models.Coords{},
		LastInteraction: false,
		Moves:           nil,
	})
	client.Hub = s.Hub
	client.Conn = conn
	client.Player = &models.Player{
		Sprite:    sprite,
		Health:    sprite.HP,
		PositionX: 0,
		PositionY: 0,
		Moves:     make(map[int]models.Coords),
	}

	go func() {
		for {
			select {
			case <-quit:
				s.Hub.Unregister <- client
				s.Hub.Broadcast <- true
			}
		}
	}()

	s.Hub.Register <- client
	return nil
}
