package handlers

import (
	"encoding/json"
	"log"
	"rogue-like/models"
	"rogue-like/services"
	"rogue-like/settings"

	"github.com/gorilla/websocket"
)

func handleKeyDown(client *models.Client, message models.WSMessage) error {
	data := models.KeyDownMessage{}
	err := json.Unmarshal(message.Data, &data)
	if err != nil {
		log.Println("error during unmarshall:", err)
		return err
	}

	switch data.Key {
	case models.ArrowLeft:
		player := &client.Player
		(*player).PositionX -= settings.MoveStep
	case models.ArrowUp:
		player := &client.Player
		(*player).PositionY -= settings.MoveStep
	case models.ArrowRight:
		player := &client.Player
		(*player).PositionX += settings.MoveStep
	case models.ArrowDown:
		player := &client.Player
		(*player).PositionY += settings.MoveStep
	}

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
	// TODO: sprite should not be hardcoded
	sprite, err := s.GetSprite(models.Templar)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	client.Hub = s.Hub
	client.Conn = conn
	client.Player = &models.Player{
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
