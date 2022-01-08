package models

type SpriteName string
type TileSet string

var (
	Warrior SpriteName = "warrior"
	Templar SpriteName = "templar"
)

var (
	// Characters TileSet = "characters"
	// Background TileSet = "background"
	Sprites TileSet = "sprites"
)

type Sprite struct {
	Name         SpriteName `json:"name"`
	TileSet      TileSet    `json:"tileSet"`
	SpriteX      int        `json:"spriteX"`
	SpriteY      int        `json:"spriteY"`
	SpriteWidth  int        `json:"spriteWidth"`
	SpriteHeight int        `json:"spriteHeight"`
	HP           int        `json:"hp"`
	MoveRange    int        `json:"moveRange"`
	AttackRange  int        `json:"attackRange"`
	XOffset      int        `json:"xOffset"`
	YOffset      int        `json:"yOffset"`
}

type Player struct {
	Sprite    Sprite `json:"sprite"`
	Health    int    `json:"health"`
	PositionX int    `json:"positionX"`
	PositionY int    `json:"positionY"`
}
