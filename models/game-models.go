package models

import (
	"errors"
	"rogue-like/helpers"
	"rogue-like/settings"
	"sort"
	"time"
)

type DropName string
type SpriteName string
type TileSet string

var (
	HealthPotion DropName = "health-potion"
)

var (
	Warrior    SpriteName = "warrior"
	Templar    SpriteName = "templar"
	Archer     SpriteName = "archer"
	Mage       SpriteName = "mage"
	MageDark   SpriteName = "mage-dark"
	Orc        SpriteName = "orc"
	OrcRed     SpriteName = "orc-red"
	OrcKing    SpriteName = "orc-king"
	SheepWhite SpriteName = "sheep-white"
	SheepGrey  SpriteName = "sheep-grey"
	SheepDark  SpriteName = "sheep-dark"
)

var (
	Sprites TileSet = "sprites"
)

var (
	ArrowLeft  = "ArrowLeft"
	ArrowUp    = "ArrowUp"
	ArrowRight = "ArrowRight"
	ArrowDown  = "ArrowDown"
	KeyA       = "a"
	KeyW       = "w"
	KeyD       = "d"
	KeyS       = "s"
)

type Animation struct {
	SpriteX      int `json:"spriteX"`
	SpriteY      int `json:"spriteY"`
	SpriteWidth  int `json:"spriteWidth"`
	SpriteHeight int `json:"spriteHeight"`
	XOffset      int `json:"xOffset"`
	YOffset      int `json:"yOffset"`
}

type DropSprite struct {
	Name         DropName                         `json:"name"`
	TileSet      TileSet                          `json:"tileSet"`
	SpriteX      int                              `json:"spriteX"`
	SpriteY      int                              `json:"spriteY"`
	SpriteWidth  int                              `json:"spriteWidth"`
	SpriteHeight int                              `json:"spriteHeight"`
	XOffset      int                              `json:"xOffset"`
	YOffset      int                              `json:"yOffset"`
	Consume      func(drop *Drop, player *Player) `json:"-"`
}

type Drop struct {
	Sprite    DropSprite `json:"sprite"`
	PositionX int        `json:"positionX"`
	PositionY int        `json:"positionY"`
	Consumed  bool       `json:"consumed"`
}

type Sprite struct {
	Name         SpriteName `json:"name"`
	TileSet      TileSet    `json:"tileSet"`
	SpriteX      int        `json:"spriteX"`
	SpriteY      int        `json:"spriteY"`
	SpriteWidth  int        `json:"spriteWidth"`
	SpriteHeight int        `json:"spriteHeight"`
	HP           int        `json:"hp"`
	// MoveRange       int        `json:"moveRange"`
	AttackRange     int       `json:"attackRange"`
	Damage          int       `json:"damage"`
	XOffset         int       `json:"xOffset"`
	YOffset         int       `json:"yOffset"`
	AnimationPeriod int       `json:"animationPeriod"`
	Animation       Animation `json:"animation"`
}

type Coords struct {
	PositionX int `json:"positionX"`
	PositionY int `json:"positionY"`
}

type Area struct {
	PosStartX int
	PosEndX   int
	PosStartY int
	PosEndY   int
}

type Player struct {
	ID               int       `json:"id"`
	Sprite           Sprite    `json:"sprite"`
	Health           int       `json:"health"`
	PositionX        int       `json:"positionX"`
	PositionY        int       `json:"positionY"`
	Dead             bool      `json:"dead"`
	Respawn          bool      `json:"-"`
	RespawnDelay     int       `json:"-"`
	RespawnPositionX int       `json:"-"`
	RespawnPositionY int       `json:"-"`
	DeathTime        time.Time `json:"-"`
}

func (player *Player) HandleMove(key string, hub *Hub) {

	_, _, err := player.ProjectMove(key, hub)
	if err != nil {
		return
	}

MakeMovement:
	for m := 0; m < settings.MoveRange; m += settings.MoveStep {
		player.Move(key)
		// for _, e := range hub.Enemies {
		// 	// TODO: create logic here
		// 	e.Move(ArrowUp)
		// }
		hub.Broadcast <- true
		time.Sleep(time.Duration(player.Sprite.AnimationPeriod) * time.Millisecond / settings.MoveRange / 4)

		overlap := 5
		if m > overlap && m < overlap+2 {
			for _, drop := range hub.Drops {
				if drop.Consumed {
					continue
				}
				if player.FoundDrop(*drop) {
					drop.Sprite.Consume(drop, player)
					// TODO: create logic to consume drops
				}
			}
		}

		if m >= overlap && !player.Dead {
		CheckOverlap:
			for _, enemy := range hub.Enemies {
				if enemy.Dead {
					continue CheckOverlap
				}
				cx, cy := player.GetCollisionsTo(*enemy, 0)
				if cx && cy {
					player.Attack(enemy)
					if enemy.Dead {
						hub.Drops = append(hub.Drops, &Drop{
							// TODO: drops should not be hardcoded
							Sprite:    *hub.DropSprites[0],
							PositionX: enemy.PositionX,
							PositionY: enemy.PositionY,
						})
					}
					for mb := overlap; mb >= 0; mb -= settings.MoveStep {
						player.Move(OppositeKey(key))
						hub.Broadcast <- true
						time.Sleep(time.Duration(player.Sprite.AnimationPeriod) * time.Millisecond / settings.MoveRange / 8)
					}
					break MakeMovement
				}
			}
		}
	}

}

func (player *Player) UpdateHP(value int) {
	player.Health += value
	if player.Health > player.Sprite.HP {
		player.Health = player.Sprite.HP
	}
	if player.Health < 0 {
		player.Health = 0
		player.Dead = true
		player.DeathTime = time.Now()
	}
}

func (player *Player) Attack(enemy *Player) {
	if enemy.Health == enemy.Sprite.HP || enemy.Health%enemy.Sprite.HP >= settings.PercentageToAttackBack {
		player.UpdateHP(-enemy.Sprite.Damage / 2)
	}
	enemy.UpdateHP(-player.Sprite.Damage)
}

func (player *Player) GetArea() Area {
	return Area{
		PosStartX: player.PositionX,
		PosEndX:   player.PositionX + player.Sprite.SpriteWidth,
		PosStartY: player.PositionY,
		PosEndY:   player.PositionY + player.Sprite.SpriteHeight,
	}
}

func (player *Player) GetViewArea() Area {
	return Area{
		PosStartX: player.PositionX - settings.ViewAreaOffsetX,
		PosEndX:   player.PositionX + settings.ViewAreaOffsetX,
		PosStartY: player.PositionY - settings.ViewAreaOffsetY,
		PosEndY:   player.PositionY + settings.ViewAreaOffsetY,
	}
}

func (player *Player) CanSee(p Player) bool {
	va := player.GetViewArea()
	return helpers.IsInsideViewArea(
		va.PosStartX,
		va.PosEndX,
		va.PosStartY,
		va.PosEndY,
		p.PositionX,
		p.PositionY,
	)
}

func (player *Player) Move(key string) {
	// TODO: return a boolean if player actually moved
	switch key {
	case ArrowLeft, KeyA:
		player.PositionX -= settings.MoveStep
	case ArrowUp, KeyW:
		player.PositionY -= settings.MoveStep
	case ArrowRight, KeyD:
		player.PositionX += settings.MoveStep
	case ArrowDown, KeyS:
		player.PositionY += settings.MoveStep
	default:
		return
	}
}

func (player *Player) ProjectMove(key string, hub *Hub) (x int, y int, err error) {
	x = player.PositionX
	y = player.PositionY
	switch key {
	case ArrowLeft, KeyA:
		x -= settings.MoveRange
	case ArrowUp, KeyW:
		y -= settings.MoveRange
	case ArrowRight, KeyD:
		x += settings.MoveRange
	case ArrowDown, KeyS:
		y += settings.MoveRange
	}

	if x < hub.MapArea.PosStartX || x > hub.MapArea.PosEndX {
		return x, y, errors.New("map limit")
	}
	if y < hub.MapArea.PosStartY || y > hub.MapArea.PosEndY {
		return x, y, errors.New("map limit")
	}

	idx := helpers.GetTileIndexByPositions(x, y, hub.FloorLayer.Width)
	_, found := hub.FloorLayer.TileMap[idx]
	if found {
		return x, y, errors.New("occupied tile")
	}

	return x, y, nil
}

func (player *Player) GetClosePlayers(players []*Player, offset int) []*Player {
	var closePlayers []*Player
	for _, p := range players {
		cx, cy := player.GetCollisionsTo(*p, offset)
		if cx && cy {
			closePlayers = append(closePlayers, p)
		}
	}
	return closePlayers
}

func (player *Player) GetClosestPlayer(players []*Player) *Player {
	distancesMap := make(map[int]*Player)
	keys := make([]int, 0, len(players))
	for _, p := range players {
		diffX := player.PositionX - p.PositionX
		diffY := player.PositionY - p.PositionY
		key := (diffX * diffX) + (diffY * diffY)
		distancesMap[key] = p
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return distancesMap[keys[0]]
}

func (player *Player) GetCollisionsTo(player2 Player, offset int) (bool, bool) {
	return helpers.HasCollision(
		player.PositionX,
		player.PositionY,
		player2.PositionX,
		player2.PositionY,
		player.Sprite.SpriteWidth+player.Sprite.XOffset,
		player.Sprite.SpriteHeight+player.Sprite.YOffset,
		player2.Sprite.SpriteWidth+player2.Sprite.XOffset,
		player2.Sprite.SpriteHeight+player2.Sprite.YOffset,
		offset,
	)
}

func (player *Player) FoundDrop(drop Drop) bool {
	cx, cy := helpers.HasCollision(
		player.PositionX,
		player.PositionY,
		drop.PositionX,
		drop.PositionY,
		player.Sprite.SpriteWidth+player.Sprite.XOffset,
		player.Sprite.SpriteHeight+player.Sprite.YOffset,
		drop.Sprite.SpriteWidth+drop.Sprite.XOffset,
		drop.Sprite.SpriteHeight+drop.Sprite.YOffset,
		0,
	)
	return cx && cy
}

func OppositeKey(key string) string {
	switch key {
	case ArrowLeft, KeyA:
		return ArrowRight
	case ArrowUp, KeyW:
		return ArrowDown
	case ArrowRight, KeyD:
		return ArrowLeft
	case ArrowDown, KeyS:
		return ArrowUp
	default:
		return ""
	}
}
