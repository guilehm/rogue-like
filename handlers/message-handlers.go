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

	if client.Player.Dead {
		return nil
	}

	if key == models.KeySpace {
		enemies := client.Hub.GetAliveEnemies(0)
		client.Player.GetClosePlayers(enemies, client.Player.Sprite.AttackRange*8)
		if len(enemies) == 0 {
			return nil
		}
		enemy := client.Player.GetClosestPlayer(enemies)
		p := client.Player.CreateProjectileTo(enemy)
		client.Hub.Projectiles[p] = true
		go client.Player.Shoot(
			enemy,
			p,
			client.Hub,
		)
		client.Hub.Broadcast <- true
	} else {
		client.Player.HandleMove(key, client.Hub)
	}
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

	client.Hub = s.Hub
	client.Conn = conn
	client.Player = &models.Player{
		ID:        data.ID,
		Sprite:    sprite,
		Health:    sprite.HP,
		PositionX: 8,
		PositionY: 8,
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
