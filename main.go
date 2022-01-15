package main

import (
	"log"
	"net/http"
	"os"
	"rogue-like/handlers"
	"rogue-like/models"
	"rogue-like/services"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("hello")

	r := mux.NewRouter()
	service := services.GameService{
		Hub: &models.Hub{
			Clients:       make(map[*models.Client]bool),
			Register:      make(chan *models.Client),
			Unregister:    make(chan *models.Client),
			Broadcast:     make(chan bool),
			PlayerSprites: []models.Sprite{},
			EnemySprites:  []models.Sprite{},
		},
	}
	service.CreateSprites()
	go service.Start()

	r.HandleFunc("/ws/rogue/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RogueLikeHandler(&service, w, r)
	})
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
