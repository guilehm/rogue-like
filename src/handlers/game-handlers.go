package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"rogue-like/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func RogueLikeHandler(hub *models.Hub, w http.ResponseWriter, r *http.Request) {
	// TODO: do not allow all origins
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	quit := make(chan bool)
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println("error during connection upgrade:", err)
		return
	}

	client := &models.Client{}
	for {
		message := models.WSMessage{}
		err = conn.ReadJSON(&message)
		if err != nil {
			if errors.Is(err.(*websocket.CloseError), err) {
				log.Println("connection closed")
				quit <- true
				return
			} else {
				log.Println("could not read message:", err)
				continue
			}
		}

		switch message.MessageType {
		case models.UserJoins:
			data := models.UserJoinsMessage{}
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				log.Println("error during unmarshall:", err)
				break
			}

			// TODO: sprite should not be hardcoded
			sprite, err := hub.GetSprite(models.Warrior)
			if err != nil {
				log.Println(err.Error())
				break
			}

			client := &models.Client{
				Hub:  hub,
				Conn: conn,
				Player: &models.Player{
					Sprite:    sprite,
					Health:    sprite.HP,
					PositionX: 0,
					PositionY: 0,
				},
			}
			hub.Register <- client
		}

		go func() {
			for {
				select {
				case <-quit:
					hub.Unregister <- client
				}
			}
		}()
	}
}
