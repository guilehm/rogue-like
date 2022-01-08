package main

import (
	"log"
	"net/http"
	"os"
	"rogue-like/handlers"
	"rogue-like/models"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Hello")

	r := mux.NewRouter()
	hub := models.Hub{
		Clients: make(map[*models.Client]bool),
	}
	go hub.Start()

	r.HandleFunc("/ws/rogue-like/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RogueLikeHandler(&hub, w, r)
	})
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
