package levels

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	WorldLevelData *LevelData
)

type LevelData struct {
	Layers   []*Layer
	TileSets []*TileSet

	Width  int
	Height int

	TileWidth  int
	TileHeight int
}

type Layer struct {
	Id int

	Data []int

	Height int
	Width  int

	Name      string
	LayerType string `json:"type"`
}

type TileSet struct {
	FirstGid int

	Image   string
	Texture rl.Texture2D

	TileHeight int
	TileWidth  int

	Columns int
}

type Tile struct {
	Texture    rl.Texture2D
	Height     int
	Width      int
	TextureRec rl.Rectangle
}
