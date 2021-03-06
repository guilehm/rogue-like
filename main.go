package main

import (
	"log"
	"net/http"
	"os"
	"rogue-like/handlers"
	"rogue-like/models"
	"rogue-like/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("hello")

	service := services.GameService{
		Hub: &models.Hub{
			Clients:           make(map[*models.Client]bool),
			Register:          make(chan *models.Client),
			Unregister:        make(chan *models.Client),
			Broadcast:         make(chan bool),
			PlayerSprites:     []models.Sprite{},
			EnemySprites:      []models.Sprite{},
			DropSprites:       []models.DropSprite{},
			ProjectileSprites: []models.ProjectileSprite{},
			Projectiles:       make(map[*models.Projectile]bool),
		},
	}
	service.Hub.LevelMap = service.GenerateGameLevelsMap()
	service.CreateSprites()
	service.CreateEnemies()
	service.CreateFloorTiles()

	go service.Start()
	go service.RespawnEnemies()
	go service.IncreasePlayersHealth()
	go service.FollowPlayers()
	go service.ClearProjectiles()

	r := mux.NewRouter()
	// TODO: do not allow all origins
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(r)
	r.HandleFunc("/sprites/", func(w http.ResponseWriter, r *http.Request) {
		handlers.SpriteList(service.Hub, w, r)
	})
	r.HandleFunc("/ws/rogue/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RogueLikeHandler(&service, w, r)
	})
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), handler)
}
