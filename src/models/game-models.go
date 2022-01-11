package models

import (
	"rogue-like/settings"
)

type SpriteName string
type TileSet string

var (
	Warrior SpriteName = "warrior"
	Templar SpriteName = "templar"
	Archer  SpriteName = "archer"
	Mage    SpriteName = "mage"
	Orc     SpriteName = "orc"
)

var (
	Sprites TileSet = "sprites"
)

var (
	ArrowLeft  = "ArrowLeft"
	ArrowUp    = "ArrowUp"
	ArrowRight = "ArrowRight"
	ArrowDown  = "ArrowDown"
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
	Name            SpriteName `json:"name"`
	TileSet         TileSet    `json:"tileSet"`
	SpriteX         int        `json:"spriteX"`
	SpriteY         int        `json:"spriteY"`
	SpriteWidth     int        `json:"spriteWidth"`
	SpriteHeight    int        `json:"spriteHeight"`
	HP              int        `json:"hp"`
	MoveRange       int        `json:"moveRange"`
	AttackRange     int        `json:"attackRange"`
	XOffset         int        `json:"xOffset"`
	YOffset         int        `json:"yOffset"`
	AnimationPeriod int        `json:"animationPeriod"`
	Animation       Animation  `json:"animation"`
}

type Coords struct {
	PositionX int `json:"positionX"`
	PositionY int `json:"positionY"`
}

type Player struct {
	Sprite    Sprite `json:"sprite"`
	Health    int    `json:"health"`
	PositionX int    `json:"positionX"`
	PositionY int    `json:"positionY"`
	// LastPosition Coords         `json:"lastPosition"`
	// Moves        map[int]Coords `json:"-"`
}

func (player *Player) Move(key string) {
	// TODO: return a boolean if player actually moved
	// player.LastPosition = player.Moves[len(player.Moves)]
	switch key {
	case ArrowLeft:
		player.PositionX -= settings.MoveStep
	case ArrowUp:
		player.PositionY -= settings.MoveStep
	case ArrowRight:
		player.PositionX += settings.MoveStep
	case ArrowDown:
		player.PositionY += settings.MoveStep
	default:
		return
	}
	// player.Moves[len(player.Moves)+1] = Coords{
	// 	PositionX: player.PositionX,
	// 	PositionY: player.PositionY,
	// }
}
