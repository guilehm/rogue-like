package handlers

import (
	"encoding/json"
	"log"
	"rogue-like/models"
	"rogue-like/services"

	"github.com/gorilla/websocket"
)

func handleUserJoins(s *services.GameService, conn *websocket.Conn, message models.WSMessage) error {

	data := models.UserJoinsMessage{}
	err := json.Unmarshal(message.Data, &data)
	if err != nil {
		log.Println("error during unmarshall:", err)
		return err
	}

	// TODO: sprite should not be hardcoded
	sprite, err := s.GetSprite(models.Warrior)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	client := &models.Client{
		Hub:  s.Hub,
		Conn: conn,
		Player: &models.Player{
			Sprite:    sprite,
			Health:    sprite.HP,
			PositionX: 0,
			PositionY: 0,
		},
	}
	s.Hub.Register <- client
}
