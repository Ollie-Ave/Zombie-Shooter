package levels

import (
	"cmp"
	"slices"
	"strings"

	"github.com/Ollie-Ave/Zombie-Shooter/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func RenderLevelData(levelData *LevelData, colliderData *LevelColliderData) {
	for layerIndex, layer := range levelData.Layers {
		if !strings.Contains(layer.Name, "Spawner_") {
			renderTileLayer(layer, levelData)
		}

		if shared.IsDebugMode() {
			renderDebugDataForLayer(layer, layerIndex, levelData, colliderData)
		}
	}
}

func renderDebugDataForLayer(layer *Layer, layerIndex int, levelData *LevelData, colliderData *LevelColliderData) {
	if strings.Contains(layer.Name, "Spawner_") {
		renderTileLayer(layer, levelData)
	}

	if strings.Contains(layer.Name, "Collidable_") {
		renderCollisionData(
			colliderData.ColliderLayers[layerIndex],
			rl.NewVector2(float32(levelData.TileWidth), float32(levelData.TileHeight)),
		)
	}
}

func renderTileLayer(layer *Layer, levelData *LevelData) {
	x := -1
	y := 0

	for _, tileId := range layer.Data {
		x, y = getNewTilePosition(x, y, layer.Width)

		if tileId == 0 {
			continue
		}

		tile := getTileData(tileId, levelData)

		tilePosition := rl.NewVector2(
			float32(x*tile.Width),
			float32(y*tile.Height))

		rl.DrawTextureRec(
			tile.Texture,
			tile.TextureRec,
			tilePosition,
			rl.White)
	}
}

func renderCollisionData(layer *ColliderLayer, tileSize rl.Vector2) {
	for x, col := range layer.Data {
		for y, value := range col {
			if value {
				xPos := int32(x * int(tileSize.X))
				yPos := int32(y * int(tileSize.Y))

				rl.DrawRectangleLines(
					xPos,
					yPos,
					int32(tileSize.X),
					int32(tileSize.Y),
					rl.Purple,
				)
			}
		}
	}
}

func getTileData(tileId int, levelData *LevelData) *Tile {
	tileSet := getTileSetById(tileId, levelData)
	tileX, tileY := getTilePositionByTileId(tileId, tileSet)

	return &Tile{
		Texture: tileSet.Texture,
		Height:  tileSet.TileHeight,
		Width:   tileSet.TileWidth,
		TextureRec: rl.NewRectangle(
			float32(tileX),
			float32(tileY),
			float32(tileSet.TileWidth),
			float32(tileSet.TileHeight)),
	}
}

func getTileSetById(id int, levelData *LevelData) *TileSet {
	var returnValue *TileSet

	slices.SortFunc(levelData.TileSets, func(a, b *TileSet) int {
		return cmp.Compare(a.FirstGid, b.FirstGid)
	})

	for _, tileSet := range levelData.TileSets {
		if tileSet.FirstGid > id {
			return returnValue
		}

		returnValue = tileSet
	}

	return returnValue
}

func getTilePositionByTileId(id int, tileSet *TileSet) (int, int) {
	x, y := 0, 0

	for i := tileSet.FirstGid; i < id; i++ {
		x++

		if x == tileSet.Columns {
			x = 0
			y++
		}
	}

	return x * tileSet.TileWidth, y * tileSet.TileHeight
}

func getNewTilePosition(x, y int, maxX int) (int, int) {
	if x < (maxX - 1) {
		x++
	} else {
		x = 0
		y++
	}

	return x, y
}
