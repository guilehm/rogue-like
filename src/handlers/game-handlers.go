package handlers

import (
	"log"
	"net/http"
	"rogue-like/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func RogueLikeHandler(hub *models.Hub, w http.ResponseWriter, r *http.Request) {
	// TODO: do not allow all origins
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println("Error during connection upgrade:", err)
		return
	}
}
