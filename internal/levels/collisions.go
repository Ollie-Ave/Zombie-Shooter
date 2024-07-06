package levels

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	WorldColliderData *LevelColliderData
)

type LevelColliderData struct {
	colliderLayers []*ColliderLayer
}

type ColliderLayer struct {
	isCollidable bool
	data         [][]bool
}

func LoadWorldColliderData(levelData *LevelData) *LevelColliderData {
	colliderData := &LevelColliderData{
		colliderLayers: make([]*ColliderLayer, len(levelData.Layers)),
	}

	for layerIndex, layer := range levelData.Layers {
		colliderData.colliderLayers[layerIndex] = loadColliderDataForLayer(
			layer,
			levelData.Width,
			levelData.Height,
		)

	}

	return colliderData
}

func HitboxCollidesWithWorld(hitbox rl.Rectangle) bool {
	for _, layer := range WorldColliderData.colliderLayers {
		if layer.isCollidable && hitboxCollidesOnThisLayer(layer, hitbox) {
			return true
		}
	}

	return false
}

func hitboxCollidesOnThisLayer(layer *ColliderLayer, hitbox rl.Rectangle) bool {
	for x, col := range layer.data {
		for y, tileIsCollidable := range col {
			if tileIsCollidable &&
				rl.CheckCollisionRecs(
					hitbox,
					rl.NewRectangle(
						float32(x*WorldLevelData.TileWidth),
						float32(y*WorldLevelData.TileHeight),
						float32(WorldLevelData.TileWidth),
						float32(WorldLevelData.TileHeight),
					),
				) {
				return true
			}
		}
	}

	return false
}

func loadColliderDataForLayer(layer *Layer, worldWidth, worldHeight int) *ColliderLayer {
	data := assign2DArrayBuffer[bool](worldWidth, worldHeight)

	for x := 0; x < worldWidth; x++ {
		for y := 0; y < worldHeight; y++ {
			layerDataIndex := y*worldWidth + x

			data[x][y] = layer.Data[layerDataIndex] != 0
		}
	}

	return &ColliderLayer{
		data:         data,
		isCollidable: strings.Contains(layer.Name, "Collidable_"),
	}
}

func assign2DArrayBuffer[T any](rows, cols int) [][]T {
	buffer := make([][]T, cols)

	for i := range buffer {
		buffer[i] = make([]T, rows)
	}

	return buffer
}
