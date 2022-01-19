package models

type TileSetData struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Layers []Layer `json:"layers"`
}

type Layer struct {
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Width   int          `json:"width"`
	Height  int          `json:"height"`
	Data    []int        `json:"data"`
	TileMap map[int]Tile `json:"tileMap"`
}
