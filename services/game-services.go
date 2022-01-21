package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"rogue-like/helpers"
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
				player.Health += 20
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
			SightDistance:   2,
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
			SightDistance:   3,
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
			SightDistance:   2,
			AnimationPeriod: 1000,
			Animation: models.Animation{
				SpriteX:      27,
				SpriteY:      19,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
		},
		{
			Name:            models.MageDark,
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
			SightDistance:   3,
			AnimationPeriod: 800,
			Animation: models.Animation{
				SpriteX:      27,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
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
			HP:              140,
			Damage:          80,
			AttackRange:     1,
			SightDistance:   4,
			AnimationPeriod: 1100,
			Animation: models.Animation{
				SpriteX:      36,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
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
			HP:              160,
			Damage:          100,
			AttackRange:     1,
			SightDistance:   4,
			AnimationPeriod: 1200,
			Animation: models.Animation{
				SpriteX:      54,
				SpriteY:      37,
				SpriteWidth:  8,
				SpriteHeight: 8,
			},
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
			HP:              180,
			Damage:          120,
			AttackRange:     1,
			SightDistance:   5,
			AnimationPeriod: 1250,
			Animation: models.Animation{
				SpriteX:      63,
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
			SightDistance:   2,
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
			SightDistance:   2,
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
			SightDistance:   2,
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
			SightDistance:   2,
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

func (s *GameService) createEnemy(name models.SpriteName, px, py, respDelay int) *models.Player {
	if px%8 != 0 || py%8 != 0 {
		log.Fatal("invalid position while creating ", name)
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

func getNextMoveKey(p1, p2 *models.Player) (key, alternative string, attack bool) {
	// p1 is following p2

	diffX := helpers.Abs(p1.PositionX - p2.PositionX)
	diffY := helpers.Abs(p1.PositionY - p2.PositionY)

	if (diffX <= 8 && diffY == 0) || (diffY <= 8 && diffX == 0) {
		attack = true
	}

	if diffX >= diffY {
		// TODO: checking >= for now. create condition for == to move X or Y
		// move X axis
		if p1.PositionY <= p2.PositionY {
			alternative = models.ArrowDown
		} else {
			alternative = models.ArrowUp
		}

		if p1.PositionX <= p2.PositionX {
			key = models.ArrowRight
		} else {
			key = models.ArrowLeft
		}
	} else {
		// move Y axis
		if p1.PositionX <= p2.PositionX {
			alternative = models.ArrowRight
		} else {
			alternative = models.ArrowLeft
		}
		if p1.PositionY <= p2.PositionY {
			key = models.ArrowDown
		} else {
			key = models.ArrowUp
		}
	}
	return key, alternative, attack
}

func (s *GameService) FollowPlayers() {
	for {
		for _, enemy := range s.Hub.Enemies {
			if enemy.Dead {
				continue
			}
			enemy := enemy
			go func() {
				var players []*models.Player
				for c := range s.Hub.Clients {
					players = append(players, c.Player)
				}
				closePlayers := enemy.GetClosePlayers(players, enemy.Sprite.SightDistance*8)
				if len(closePlayers) > 0 {
					closestPlayer := enemy.GetClosestPlayer(closePlayers)
					diffX := helpers.Abs(enemy.PositionX - closestPlayer.PositionX)
					diffY := helpers.Abs(enemy.PositionY - closestPlayer.PositionY)

					if (diffX <= 8 && diffY == 0) || (diffY <= 8 && diffX == 0) {
						// TODO: attack!!!!!!!!
						return
					}
					if diffX > diffY {
						// move X axis
						// TODO: handle errors to choose another position
						if enemy.PositionX <= closestPlayer.PositionX {
							err := enemy.ProjectAndMove(models.ArrowRight, s.Hub)
							if err != nil {
								return
							}
						} else {
							err := enemy.ProjectAndMove(models.ArrowLeft, s.Hub)
							if err != nil {
								return
							}
						}
					} else if diffX == diffY {
						// move X or Y axis
					} else {
						// move Y axis
						if enemy.PositionY <= closestPlayer.PositionY {
							err := enemy.ProjectAndMove(models.ArrowDown, s.Hub)
							if err != nil {
								return
							}
						} else {
							err := enemy.ProjectAndMove(models.ArrowUp, s.Hub)
							if err != nil {
								return
							}
						}
					}
					s.Hub.Broadcast <- true
				}
			}()
		}
		// TODO: this time should be as low as possible
		// monsters may have different time delay to move
		time.Sleep(800 * time.Millisecond)
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
					Type:    models.Broadcast,
					Players: filteredPlayers,
					Enemies: filteredEnemies,
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
