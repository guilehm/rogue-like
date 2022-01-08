package models

type Sprite struct {
	Image        string `json:"image"`
	SpriteX      int    `json:"spriteX"`
	SpriteY      int    `json:"spriteY"`
	SpriteWidth  int    `json:"spriteWidth"`
	SpriteHeight int    `json:"spriteHeight"`
	HP           int    `json:"hp"`
	MoveRange    int    `json:"moveRange"`
	AttackRange  int    `json:"attackRange"`
}
