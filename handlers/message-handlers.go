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

	if client.Player.Dead {
		return nil
	}
MakeMovement:
	for m := 0; m < settings.MoveRange; m += settings.MoveStep {
		client.Player.Move(key)
		// for _, e := range client.Hub.Enemies {
		// 	// TODO: create logic here
		// 	e.Move(models.ArrowUp)
		// }
		client.Hub.Broadcast <- true
		time.Sleep(time.Duration(client.Player.Sprite.AnimationPeriod) * time.Millisecond / settings.MoveRange / 4)

		overlap := 5
		if m > overlap && m < overlap+2 {
			for _, drop := range client.Hub.Drops {
				if drop.Consumed {
					continue
				}
				if client.Player.FoundDrop(*drop) {
					drop.Sprite.Consume(drop, client.Player)
					// TODO: create logic to consume drops
				}
			}
		}

		if m >= overlap && !client.Player.Dead {
		CheckOverlap:
			for _, enemy := range client.Hub.Enemies {
				if enemy.Dead {
					continue CheckOverlap
				}
				cx, cy := client.Player.GetCollisionsTo(*enemy, 0)
				if cx && cy {
					client.Player.Attack(enemy)
					if enemy.Dead {
						client.Hub.Drops = append(client.Hub.Drops, &models.Drop{
							// TODO: Drops should not be hardcoded
							Sprite:    *client.Hub.DropSprites[0],
							PositionX: enemy.PositionX,
							PositionY: enemy.PositionY,
						})
					}
					for mb := overlap; mb >= 0; mb -= settings.MoveStep {
						client.Player.Move(models.OppositeKey(key))
						client.Hub.Broadcast <- true
						time.Sleep(time.Duration(client.Player.Sprite.AnimationPeriod) * time.Millisecond / settings.MoveRange / 8)
					}
					break MakeMovement
				}
			}
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
	sprite := s.Hub.PlayerSprites[rand.Int()%len(s.Hub.PlayerSprites)]

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
