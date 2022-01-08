package handlers

import (
	"errors"
	"fmt"
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
		log.Println("Error during connection upgrade:", err)
		return
	}

	for {
		message := models.WSMessage{}
		err = conn.ReadJSON(&message)
		if err != nil {
			if errors.Is(err.(*websocket.CloseError), err) {
				log.Println("Connection closed")
				quit <- true
				return
			} else {
				fmt.Println("Could not read message:", err)
				continue
			}
		}
	}
}
