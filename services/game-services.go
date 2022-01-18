package services

import (
	"errors"
	"log"
	"rogue-like/models"
	"rogue-like/settings"
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
			Name:            models.OrcRed,
			TileSet:         models.Sprites,
			SpriteX:         36,
			SpriteY:         27,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              110,
			Damage:          50,
			AttackRange:     1,
			AnimationPeriod: 800,
			Animation: models.Animation{
				SpriteX:      36,
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
			HP:              70,
			Damage:          65,
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

func (s *GameService) CreateEnemies() {
	orcSprite, _ := s.GetSprite(models.Orc, "enemy")
	darkMageSprite, _ := s.GetSprite(models.DarkMage, "enemy")
	orcKingSprite, _ := s.GetSprite(models.OrcKing, "enemy")
	orcRedSprite, _ := s.GetSprite(models.OrcRed, "enemy")
	s.Hub.Enemies = append(s.Hub.Enemies,
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcSprite,
			Health:           orcSprite.HP,
			PositionX:        184,
			PositionY:        24,
			Respawn:          true,
			RespawnPositionX: 184,
			RespawnPositionY: 24,
			RespawnDelay:     30,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcSprite,
			Health:           orcSprite.HP,
			PositionX:        192,
			PositionY:        72,
			Respawn:          true,
			RespawnPositionX: 192,
			RespawnPositionY: 72,
			RespawnDelay:     30,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcSprite,
			Health:           orcSprite.HP,
			PositionX:        128,
			PositionY:        40,
			Respawn:          true,
			RespawnPositionX: 128,
			RespawnPositionY: 40,
			RespawnDelay:     30,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcSprite,
			Health:           orcSprite.HP,
			PositionX:        184,
			PositionY:        96,
			Respawn:          true,
			RespawnPositionX: 184,
			RespawnPositionY: 96,
			RespawnDelay:     30,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcSprite,
			Health:           orcSprite.HP,
			PositionX:        88,
			PositionY:        64,
			Respawn:          true,
			RespawnPositionX: 88,
			RespawnPositionY: 64,
			RespawnDelay:     30,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcRedSprite,
			Health:           orcRedSprite.HP,
			PositionX:        128,
			PositionY:        104,
			Respawn:          true,
			RespawnPositionX: 128,
			RespawnPositionY: 104,
			RespawnDelay:     45,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           darkMageSprite,
			Health:           darkMageSprite.HP,
			PositionX:        240,
			PositionY:        96,
			Respawn:          true,
			RespawnPositionX: 240,
			RespawnPositionY: 96,
			RespawnDelay:     60,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcKingSprite,
			Health:           orcKingSprite.HP,
			PositionX:        128,
			PositionY:        136,
			Respawn:          true,
			RespawnPositionX: 128,
			RespawnPositionY: 136,
			RespawnDelay:     60,
		},
		&models.Player{
			ID:               int(time.Now().UnixNano()),
			Sprite:           orcKingSprite,
			Health:           orcKingSprite.HP,
			PositionX:        128 + 8,
			PositionY:        136,
			Respawn:          true,
			RespawnPositionX: 128 + 8,
			RespawnPositionY: 136,
			RespawnDelay:     60,
		},
	)
}

func (s *GameService) IncreasePlayersHealth() {
	for {
		for client := range s.Hub.Clients {
			if client.Player.Dead || client.Player.Health >= client.Player.Sprite.HP {
				continue
			}
			client.Player.UpdateHP(settings.IncreasePlayersHealthValue)
		}
		s.Hub.Broadcast <- true
		time.Sleep(settings.IncreasePlayersHealthCheckTime)
	}
}

func (s *GameService) RespawnEnemies() {
	for {
		for _, enemy := range s.Hub.Enemies {
			if !enemy.Respawn || !enemy.Dead {
				continue
			}
			now := time.Now()
			if enemy.DeathTime.Add(time.Duration(enemy.RespawnDelay) * time.Second).Before(now) {
				// TODO: check area if it has no collision
				enemy.Dead = false
				enemy.Health = enemy.Sprite.HP
				enemy.PositionX = enemy.RespawnPositionX
				enemy.PositionY = enemy.RespawnPositionY
				s.Hub.Broadcast <- true
			}
		}
		time.Sleep(settings.RespawnCheckTime)
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
