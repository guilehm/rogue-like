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

type Tile struct {
	Index  int  `json:"index"`
	Row    int  `json:"row"`
	Column int  `json:"column"`
	Area   Area `json:"area"`
	Value  int  `json:"value"`
}

func (t Tile) GetTileArea() Area {
	return Area{
		PosStartX: t.Column * 8,
		PosEndX:   t.Column*8 + 8,
		PosStartY: t.Row * 8,
		PosEndY:   t.Row*8 + 8,
	}
}
