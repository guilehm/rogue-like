package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
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

func (s *GameService) GetProjectileSprite(name models.ProjectileName) (models.ProjectileSprite, error) {
	for _, sprite := range s.Hub.ProjectileSprites {
		if sprite.Name == name {
			return sprite, nil
		}
	}
	return models.ProjectileSprite{}, errors.New("sprite not found")

}

func (s *GameService) CreateSprites() {
	bolt := models.ProjectileSprite{
		Name:         models.Bolt,
		TileSet:      models.Sprites,
		SpriteX:      73,
		SpriteY:      91,
		SpriteWidth:  6,
		SpriteHeight: 1,
		XOffset:      0,
		YOffset:      4,
	}
	s.Hub.ProjectileSprites = []models.ProjectileSprite{
		bolt,
	}
	s.Hub.DropSprites = []models.DropSprite{
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
				player.UpdateHP(25)
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
			AttackRange:     0,
			SightDistance:   3,
			AnimationPeriod: 650,
			Animation: models.Animation{
				SpriteX:      0,
				SpriteY:      19,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 1800,
			MoveTimeCooldown:   700,
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
			AttackRange:     0,
			SightDistance:   3,
			AnimationPeriod: 800,
			Animation: models.Animation{
				SpriteX:      36,
				SpriteY:      19,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 2000,
			MoveTimeCooldown:   600,
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
			AttackRange:     0,
			SightDistance:   3,
			AnimationPeriod: 1000,
			Animation: models.Animation{
				SpriteX:      27,
				SpriteY:      19,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 2000,
			MoveTimeCooldown:   1000,
		},
		{
			Name:         models.MageDark,
			TileSet:      models.Sprites,
			SpriteX:      27,
			SpriteY:      45,
			SpriteWidth:  8,
			SpriteHeight: 9,
			XOffset:      0,
			YOffset:      -1,
			HP:           80,
			Damage:       50,
			AttackRange:  2,
			// TODO: mage should not shoot bolt
			ProjectileSprite: bolt,
			SightDistance:    4,
			AnimationPeriod:  800,
			Animation: models.Animation{
				SpriteX:      27,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 4000,
			MoveTimeCooldown:   1200,
		},
		{
			Name:            models.SheepWhite,
			TileSet:         models.Sprites,
			SpriteX:         36,
			SpriteY:         45,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              180,
			Damage:          80,
			AttackRange:     0,
			SightDistance:   4,
			AnimationPeriod: 1100,
			Animation: models.Animation{
				SpriteX:      36,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 2100,
			MoveTimeCooldown:   500,
		},
		{
			Name:            models.SheepGrey,
			TileSet:         models.Sprites,
			SpriteX:         54,
			SpriteY:         45,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              210,
			Damage:          100,
			AttackRange:     0,
			SightDistance:   4,
			AnimationPeriod: 1200,
			Animation: models.Animation{
				SpriteX:      54,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 2000,
			MoveTimeCooldown:   450,
		},
		{
			Name:            models.SheepDark,
			TileSet:         models.Sprites,
			SpriteX:         63,
			SpriteY:         45,
			SpriteWidth:     8,
			SpriteHeight:    9,
			XOffset:         0,
			YOffset:         -1,
			HP:              240,
			Damage:          120,
			AttackRange:     0,
			SightDistance:   5,
			AnimationPeriod: 1250,
			Animation: models.Animation{
				SpriteX:      63,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 1800,
			MoveTimeCooldown:   400,
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
			HP:              180,
			Damage:          30,
			AttackRange:     0,
			SightDistance:   2,
			AnimationPeriod: 800,
			Animation: models.Animation{
				SpriteX:      63,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 1250,
			BonusByLevel: struct {
				HP     int
				Damage int
			}{HP: 12, Damage: 1},
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
			HP:              160,
			Damage:          40,
			AttackRange:     0,
			SightDistance:   2,
			AnimationPeriod: 1000,
			Animation: models.Animation{
				SpriteX:      54,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 1300,
			BonusByLevel: struct {
				HP     int
				Damage int
			}{HP: 10, Damage: 2},
		},
		{
			Name:             models.Archer,
			TileSet:          models.Sprites,
			SpriteX:          18,
			SpriteY:          10,
			SpriteWidth:      8,
			SpriteHeight:     8,
			XOffset:          0,
			YOffset:          0,
			HP:               60,
			Damage:           45,
			AttackRange:      3,
			ProjectileSprite: bolt,
			SightDistance:    2,
			AnimationPeriod:  600,
			Animation: models.Animation{
				SpriteX:      18,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 2000,
			BonusByLevel: struct {
				HP     int
				Damage int
			}{HP: 5, Damage: 5},
		},
		{
			Name:         models.Mage,
			TileSet:      models.Sprites,
			SpriteX:      45,
			SpriteY:      9,
			SpriteWidth:  8,
			SpriteHeight: 9,
			XOffset:      0,
			YOffset:      -1,
			HP:           55,
			Damage:       60,
			AttackRange:  2,
			// TODO: mages should not shoot bolt
			ProjectileSprite: bolt,
			SightDistance:    2,
			AnimationPeriod:  750,
			Animation: models.Animation{
				SpriteX:      45,
				SpriteY:      1,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
			AttackTimeCooldown: 2000,
			BonusByLevel: struct {
				HP     int
				Damage int
			}{HP: 4, Damage: 6},
		},
	}
}

func (s *GameService) createEnemy(name models.SpriteName, px, py, respDelay int) *models.Player {
	if px%8 != 0 || py%8 != 0 {
		log.Fatal(fmt.Sprintf("invalid position while creating %v at x:%v y:%v\n", name, px, py))
	}
	sprite, _ := s.GetSprite(name, "enemy")
	return &models.Player{
		ID:               int(time.Now().UnixNano()),
		Sprite:           sprite,
		Health:           sprite.HP,
		PositionX:        px,
		PositionY:        py,
		Dead:             false,
		Respawn:          true,
		RespawnDelay:     respDelay,
		RespawnPositionX: px,
		RespawnPositionY: py,
	}
}

func (s *GameService) CreateEnemies() {
	s.Hub.Enemies = append(s.Hub.Enemies,
		s.createEnemy(models.Orc, 184, 24, 20),
		s.createEnemy(models.Orc, 192, 72, 20),
		s.createEnemy(models.Orc, 128, 40, 20),
		s.createEnemy(models.Orc, 288, 32, 20),
		s.createEnemy(models.Orc, 184, 96, 20),
		s.createEnemy(models.OrcRed, 128, 104, 45),
		s.createEnemy(models.MageDark, 240, 96, 60),
		s.createEnemy(models.MageDark, 72, 96, 60),
		s.createEnemy(models.OrcKing, 128, 136, 60),
		s.createEnemy(models.OrcKing, 128+8, 136, 60),
		s.createEnemy(models.MageDark, 96, 168, 60),
		s.createEnemy(models.OrcRed, 64, 152, 45),
		s.createEnemy(models.OrcKing, 8, 176, 60),
		s.createEnemy(models.MageDark, 120, 240, 60),
		s.createEnemy(models.MageDark, 144, 240, 60),
		s.createEnemy(models.Orc, 168, 264, 20),
		s.createEnemy(models.OrcRed, 184, 288, 45),
		s.createEnemy(models.OrcRed, 216, 304, 45),
		s.createEnemy(models.MageDark, 232, 272, 60),
		s.createEnemy(models.OrcKing, 256, 280, 60),
		s.createEnemy(models.MageDark, 64, 336, 60),
		s.createEnemy(models.MageDark, 32, 272, 60),
		s.createEnemy(models.MageDark, 40, 352, 60),
		s.createEnemy(models.SheepWhite, 32, 352, 50),
		s.createEnemy(models.OrcRed, 80, 272, 45),
		s.createEnemy(models.OrcKing, 296, 224, 60),
		s.createEnemy(models.Orc, 200, 176, 20),
		s.createEnemy(models.Orc, 224, 184, 20),
		s.createEnemy(models.Orc, 192, 152, 20),
		s.createEnemy(models.Orc, 264, 176, 20),
		s.createEnemy(models.OrcKing, 296, 160, 60),
		s.createEnemy(models.SheepGrey, 352, 24, 30),
		s.createEnemy(models.OrcRed, 336, 72, 45),
		s.createEnemy(models.OrcRed, 376, 88, 45),
		s.createEnemy(models.SheepWhite, 320, 120, 50),
		s.createEnemy(models.SheepGrey, 384, 152, 50),
		s.createEnemy(models.SheepWhite, 256, 200, 50),
		s.createEnemy(models.SheepWhite, 120, 312, 50),
		s.createEnemy(models.SheepGrey, 344, 376, 50),
		s.createEnemy(models.Orc, 384, 352, 20),
		s.createEnemy(models.MageDark, 416, 304, 60),
		s.createEnemy(models.MageDark, 432, 320, 60),
		s.createEnemy(models.MageDark, 424, 336, 60),
		s.createEnemy(models.MageDark, 160, 400, 60),
		s.createEnemy(models.OrcKing, 144, 392, 60),
		s.createEnemy(models.SheepWhite, 104, 376, 50),
		s.createEnemy(models.SheepWhite, 216, 480, 50),
		s.createEnemy(models.MageDark, 184, 488, 60),
		s.createEnemy(models.SheepGrey, 248, 488, 50),
		s.createEnemy(models.SheepGrey, 280, 520, 50),
		s.createEnemy(models.MageDark, 248, 528, 60),
		s.createEnemy(models.MageDark, 200, 520, 60),
		s.createEnemy(models.SheepGrey, 136, 512, 50),
		s.createEnemy(models.Orc, 136, 424, 20),
		s.createEnemy(models.Orc, 184, 448, 20),
		s.createEnemy(models.Orc, 240, 440, 20),
		s.createEnemy(models.OrcRed, 288, 464, 45),
		s.createEnemy(models.OrcKing, 144, 392, 60),
		s.createEnemy(models.OrcKing, 336, 448, 60),
		s.createEnemy(models.OrcKing, 336, 464, 60),
		s.createEnemy(models.MageDark, 320, 488, 60),
		s.createEnemy(models.SheepGrey, 48, 536, 50),
		s.createEnemy(models.SheepDark, 56, 592, 40),
		s.createEnemy(models.Orc, 64, 624, 20),
		s.createEnemy(models.Orc, 80, 640, 20),
		s.createEnemy(models.Orc, 104, 664, 20),
		s.createEnemy(models.MageDark, 56, 680, 60),
		s.createEnemy(models.SheepDark, 88, 728, 40),
		s.createEnemy(models.SheepDark, 192, 728, 40), // new monster here
		s.createEnemy(models.SheepDark, 368, 728, 40), // new monster here
		s.createEnemy(models.SheepDark, 368, 648, 40), // new monster here
		s.createEnemy(models.MageDark, 336, 640, 60),
		s.createEnemy(models.MageDark, 352, 608, 60),
		s.createEnemy(models.MageDark, 408, 592, 60),
		s.createEnemy(models.OrcRed, 416, 584, 45),
		s.createEnemy(models.OrcRed, 432, 608, 45),
		s.createEnemy(models.OrcKing, 408, 544, 60),
		s.createEnemy(models.SheepWhite, 104, 376, 50),
		s.createEnemy(models.SheepWhite, 264, 600, 50),
		s.createEnemy(models.SheepWhite, 224, 600, 50),
		s.createEnemy(models.SheepDark, 192, 608, 40),
		s.createEnemy(models.SheepDark, 296, 608, 40),
		s.createEnemy(models.SheepGrey, 240, 624, 50),
		s.createEnemy(models.SheepGrey, 248, 624, 50),
		s.createEnemy(models.SheepDark, 192, 672, 40),
		s.createEnemy(models.SheepDark, 296, 672, 40),
		s.createEnemy(models.SheepDark, 216, 664, 40),
		s.createEnemy(models.SheepDark, 272, 664, 40),
	)
}

func (s *GameService) IncreasePlayersHealth() {
	for {
		for client := range s.Hub.Clients {
			if client.Player.Dead || client.Player.Health >= client.Player.GetMaxHP() {
				continue
			}
			client.Player.UpdateHP(settings.IncreasePlayersHealthValue + client.Player.Sprite.BonusByLevel.HP)
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
				// add a cooldown for respawned enemies
				enemy.LastMoveTime = now.Add(settings.CooldownForRespawnedEnemies)
				enemy.LastAttackTime = now.Add(settings.CooldownForRespawnedEnemies)
				s.Hub.Broadcast <- true
			}
		}
		time.Sleep(settings.RespawnCheckTime)
	}
}

func (s *GameService) FollowPlayers() {
	for {
		// TODO: Only needs to follow enemies who has players close to it
		for _, enemy := range s.Hub.GetAliveEnemies(0) {
			if !enemy.CanMove() {
				continue
			}
			enemy := enemy
			go func() {
				players := s.Hub.GetAlivePlayers(0)
				closePlayers := enemy.GetClosePlayers(players, enemy.Sprite.SightDistance*8)
				if len(closePlayers) > 0 {
					closestPlayer := enemy.GetClosestPlayer(closePlayers)
					key, alternative, attack := enemy.GetNextMoveKey(closestPlayer)

					if enemy.CanShoot() {
						err := enemy.HandleShoot(s.Hub, s.Hub.GetAlivePlayers(0))
						if err == nil {
							return
						}
					}

					var (
						opposite1 string
						opposite2 string
					)
					keys := []string{models.ArrowLeft, models.ArrowUp, models.ArrowRight, models.ArrowDown}
					rand.Seed(time.Now().UnixNano())
					rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })
					for _, k := range keys {
						if k == key || k == alternative || k == opposite1 {
							continue
						}
						if opposite1 == "" {
							opposite1 = k
							continue
						}
						opposite2 = k
					}

					nextMoveKey := key
					if attack {
						if !enemy.CanAttack() {
							return
						}
						enemy.MoveAndAttack(closestPlayer, "", s.Hub)
					} else {
						enemies := s.Hub.GetAliveEnemies(enemy.ID)
						x, y, err := enemy.ProjectMove(nextMoveKey, s.Hub)
						if err == nil {
							collision, _ := enemy.HasProjectedCollision(enemies, x, y)
							if collision {
								err = errors.New("collision")
							}
						}
						if err != nil {
							x, y, err = enemy.ProjectMove(alternative, s.Hub)
							if err == nil {
								collision, _ := enemy.HasProjectedCollision(enemies, x, y)
								if collision {
									err = errors.New("collision")
								}
							}
							nextMoveKey = alternative
							if err != nil {
								x, y, err = enemy.ProjectMove(opposite1, s.Hub)
								if err == nil {
									collision, _ := enemy.HasProjectedCollision(enemies, x, y)
									if collision {
										err = errors.New("collision")
									}
								}
								nextMoveKey = opposite1
								if err != nil {
									x, y, err = enemy.ProjectMove(opposite2, s.Hub)
									if err == nil {
										collision, _ := enemy.HasProjectedCollision(enemies, x, y)
										if collision {
											err = errors.New("collision")
										}
									}
									nextMoveKey = opposite2
									if err != nil {
										log.Println("could not find a good route")
										return
									}
								}
							}
						}
						err = enemy.ProjectAndMove(nextMoveKey, s.Hub)
						if err != nil {
							log.Println("Error while moving ", err)
							return
						}
					}
				}
			}()
		}
		time.Sleep(settings.FollowCheckTime)
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
				client.Player.Level = client.Player.GetLevel()
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

			var projectiles []models.Projectile
			for p := range s.Hub.Projectiles {
				projectiles = append(projectiles, *p)
			}

			for client := range s.Hub.Clients {
				// here we filter enemies and players
				// to decrease the data sent to the frontend
				var filteredEnemies []models.Player
				for _, enemy := range enemies {
					if client.Player.CanSee(enemy) {
						filteredEnemies = append(filteredEnemies, enemy)
					}
				}

				var filteredPlayers []models.Player
				for _, player := range players {
					if client.Player.CanSee(player) {
						filteredPlayers = append(filteredPlayers, player)
					}
				}

				err := client.Conn.WriteJSON(models.BroadcastMessage{
					Type:        models.Broadcast,
					Players:     filteredPlayers,
					Enemies:     filteredEnemies,
					Drops:       drops,
					Projectiles: projectiles,
				})
				if err != nil {
					log.Println("could not send message:", err)
					continue
				}
			}
		}
	}
}

func (s *GameService) CreateFloorTiles() {

	endpoint := os.Getenv("TILE_MAP_DATA_ENDPOINT")
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal("could not request tile map")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("could not read tile map response body")
	}
	var tsd models.TileSetData
	err = json.Unmarshal(body, &tsd)
	if err != nil {
		log.Fatal("could not read tile map")
	}

	var floor models.Layer
	var base models.Layer
	for _, layer := range tsd.Layers {
		if layer.Name == "floor" {
			floor = layer
		} else if layer.Name == "base" {
			base = layer
		}
	}

	floor.TileMap = make(map[int]models.Tile)
	base.TileMap = make(map[int]models.Tile)

	for index, value := range floor.Data {
		if value != 0 {
			floor.TileMap[index] = floor.CreateTile(index, value)
		}
	}

	keys := make([]int, 0, len(base.Data))
	for index, value := range base.Data {
		base.TileMap[index] = base.CreateTile(index, value)
		keys = append(keys, index)
	}

	hk := keys[len(keys)-1]
	mapArea := models.Area{
		PosStartX: 0,
		PosEndX:   base.Width*8 - 8,
		PosStartY: 0,
		PosEndY:   (hk*8)/base.Width - 8,
	}
	s.Hub.FloorLayer = floor
	s.Hub.MapArea = mapArea
}
