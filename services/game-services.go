package services

import (
	"errors"
	"log"
	"rogue-like/models"
	"time"
)

type GameService struct {
	Hub *models.Hub
}

func (s *GameService) GetSprite(name models.SpriteName, kind string) (models.Sprite, error) {
	if kind == "player" {
		for _, sprite := range s.Hub.PlayerSprites {
			if sprite.Name == name {
				return sprite, nil
			}
		}
	}
	if kind == "enemy" {
		for _, sprite := range s.Hub.EnemySprites {
			if sprite.Name == name {
				return sprite, nil
			}
		}
	}
	return models.Sprite{}, errors.New("sprite not found")
}

func (s *GameService) CreateSprites() {
	s.Hub.DropSprites = []*models.DropSprite{
		{
			Name:         models.HealthPotion,
			TileSet:      models.Sprites,
			SpriteX:      98,
			SpriteY:      90,
			SpriteWidth:  4,
			SpriteHeight: 5,
			XOffset:      2,
			YOffset:      2,
			Consume: func(drop *models.Drop, player *models.Player) {
				drop.Consumed = true
				player.Health += 10
				if player.Health >= player.Sprite.HP {
					player.Health = player.Sprite.HP
				}
			},
		},
	}
	s.Hub.EnemySprites = []models.Sprite{
		{
			Name:            models.Orc,
			TileSet:         models.Sprites,
			SpriteX:         0,
			SpriteY:         27,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              100,
			Damage:          40,
			AttackRange:     1,
			AnimationPeriod: 650,
			Animation: models.Animation{
				SpriteX:      0,
				SpriteY:      19,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
		{
			Name:            models.OrcKing,
			TileSet:         models.Sprites,
			SpriteX:         27,
			SpriteY:         27,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              120,
			Damage:          60,
			AttackRange:     1,
			AnimationPeriod: 1000,
			Animation: models.Animation{
				SpriteX:      27,
				SpriteY:      19,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
		{
			Name:            models.DarkMage,
			TileSet:         models.Sprites,
			SpriteX:         27,
			SpriteY:         45,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              80,
			Damage:          70,
			AttackRange:     1,
			AnimationPeriod: 800,
			Animation: models.Animation{
				SpriteX:      27,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
	}
	s.Hub.PlayerSprites = []models.Sprite{
		{
			Name:            models.Warrior,
			TileSet:         models.Sprites,
			SpriteX:         63,
			SpriteY:         9,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              140,
			Damage:          35,
			AttackRange:     1,
			AnimationPeriod: 800,
			Animation: models.Animation{
				SpriteX:      63,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
		{
			Name:            models.Templar,
			TileSet:         models.Sprites,
			SpriteX:         54,
			SpriteY:         9,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              100,
			Damage:          45,
			AttackRange:     1,
			AnimationPeriod: 1000,
			Animation: models.Animation{
				SpriteX:      54,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
		{
			Name:            models.Archer,
			TileSet:         models.Sprites,
			SpriteX:         18,
			SpriteY:         10,
			SpriteWidth:     8,
			SpriteHeight:    8,
			XOffset:         0,
			YOffset:         0,
			HP:              60,
			Damage:          65,
			AttackRange:     1,
			AnimationPeriod: 600,
			Animation: models.Animation{
				SpriteX:      18,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
		{
			Name:            models.Mage,
			TileSet:         models.Sprites,
			SpriteX:         45,
			SpriteY:         9,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              50,
			Damage:          75,
			AttackRange:     1,
			AnimationPeriod: 750,
			Animation: models.Animation{
				SpriteX:      45,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
	}
}

func (s *GameService) spawnEnemies() {
	// TODO: create logic to spawn enemies
	var areas []models.Area
	for _, enemy := range s.Hub.Enemies {
		areas = append(areas, enemy.GetArea())
	}
}

func (s *GameService) Start() {
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
			var players []models.Player
			for client := range s.Hub.Clients {
				players = append(players, *client.Player)
			}

			var enemies []models.Player
			for _, enemy := range s.Hub.Enemies {
				if enemy.Dead {
					continue
				}
				enemies = append(enemies, *enemy)
			}

			var drops []models.Drop
			for _, drop := range s.Hub.Drops {
				if drop.Consumed {
					continue
				}
				drops = append(drops, *drop)
			}

			for client := range s.Hub.Clients {
				err := client.Conn.WriteJSON(models.BroadcastMessage{
					Type:    models.Broadcast,
					Players: players,
					Enemies: enemies,
					Drops:   drops,
				})
				if err != nil {
					log.Println("could not send message:", err)
					continue
				}
			}
		}
	}
}
