package handlers

import (
	"encoding/json"
	"errors"
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
		return errors.New("player is dead")
	}

	if key == models.KeySpace {
		enemies := client.Hub.GetAliveEnemies(0)
		err := client.Player.HandleShoot(client.Hub, enemies)
		if err != nil {
			return err
		}
	} else {
		err := client.Player.HandleMove(key, client.Hub)
		if err != nil {
			return err
		}
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
	sprite, _ := s.GetSprite(data.Sprite, "player")

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
