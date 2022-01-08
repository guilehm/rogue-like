package services

import (
	"errors"
	"log"
	"rogue-like/models"
)

type GameService struct {
	Hub *models.Hub
}

func (s *GameService) GetSprite(name models.SpriteName) (models.Sprite, error) {
	for _, sprite := range s.Hub.Sprites {
		if sprite.Name == name {
			return sprite, nil
		}
	}
	return models.Sprite{}, errors.New("sprite not found")
}

func (s *GameService) createSprites() {
	s.Hub.Sprites = []models.Sprite{
		{
			Name:         models.Warrior,
			TileSet:      models.Sprites,
			SpriteX:      63,
			SpriteY:      9,
			SpriteWidth:  8,
			SpriteHeight: 9,
			HP:           100,
			MoveRange:    1,
			AttackRange:  1,
			XOffset:      0,
			YOffset:      -1,
		},
		{
			Name:         models.Templar,
			TileSet:      models.Sprites,
			SpriteX:      54,
			SpriteY:      9,
			SpriteWidth:  8,
			SpriteHeight: 9,
			HP:           100,
			MoveRange:    1,
			AttackRange:  1,
			XOffset:      0,
			YOffset:      -1,
		},
	}
}

func (s *GameService) Start() {
	s.createSprites()
	for {
		select {
		case client := <-s.Hub.Register:
			log.Println("registering client")
			s.Hub.Clients[client] = true
		case client := <-s.Hub.Unregister:
			if _, ok := s.Hub.Clients[client]; ok {
				log.Println("unregistering client")
				delete(s.Hub.Clients, client)
			}
		case <-s.Hub.Broadcast:
			log.Println("broadcasting")

			var players []*models.Player
			for client := range s.Hub.Clients {
				players = append(players, client.Player)
			}

			for client := range s.Hub.Clients {
				err := client.Conn.WriteJSON(models.BroadcastMessage{
					Type:    models.Broadcast,
					Players: players,
				})
				if err != nil {
					log.Println("could not send message:", err)
					continue
				}
			}
		}
	}
}
