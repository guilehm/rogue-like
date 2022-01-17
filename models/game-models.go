package models

import (
	"rogue-like/helpers"
	"rogue-like/settings"
)

type SpriteName string
type TileSet string

var (
	Warrior  SpriteName = "warrior"
	Templar  SpriteName = "templar"
	Archer   SpriteName = "archer"
	Mage     SpriteName = "mage"
	Orc      SpriteName = "orc"
	OrcKing  SpriteName = "orc-king"
	DarkMage SpriteName = "dark-mage"
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
	ID        int    `json:"id"`
	Sprite    Sprite `json:"sprite"`
	Health    int    `json:"health"`
	PositionX int    `json:"positionX"`
	PositionY int    `json:"positionY"`
	Dead      bool   `json:"dead"`
}

func (player *Player) Attack(enemy *Player) {
	enemy.Health -= player.Sprite.Damage
	player.Health -= enemy.Sprite.Damage / 2
	if enemy.Health <= 0 {
		enemy.Health = 0
		enemy.Dead = true
	}
	if player.Health <= 0 {
		player.Health = 0
		player.Dead = true
	}
}

func (player *Player) GetArea() Area {
	return Area{
		PosStartX: player.PositionX,
		PosEndX:   player.PositionX + player.Sprite.SpriteWidth,
		PosStartY: player.PositionY,
		PosEndY:   player.PositionY + player.Sprite.SpriteHeight,
	}
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
