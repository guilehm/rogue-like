package handlers

import (
	"encoding/json"
	"log"
	"math/rand"
	"rogue-like/models"
	"rogue-like/services"
	"rogue-like/settings"
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

	for m := 0; m < settings.MoveRange; m += settings.MoveStep {
		client.Player.Move(key)
		// for _, e := range client.Hub.Enemies {
		// 	// TODO: create logic here
		// 	e.Move(models.ArrowUp)
		// }
		client.Hub.Broadcast <- true
		time.Sleep(time.Duration(client.Player.Sprite.AnimationPeriod) * time.Millisecond / settings.MoveRange / 4)
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
	enemySprite := s.Hub.EnemySprites[rand.Int()%len(s.Hub.EnemySprites)]

	pX := rand.Intn(10) * 8
	s.Hub.Enemies = append(s.Hub.Enemies, &models.Player{
		Sprite:    enemySprite,
		Health:    enemySprite.HP,
		PositionX: pX,
		PositionY: 8 * 4,
	})
	client.Hub = s.Hub
	client.Conn = conn
	client.Player = &models.Player{
		ID:        data.ID,
		Sprite:    sprite,
		Health:    sprite.HP,
		PositionX: 0,
		PositionY: 0,
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
