package levels

import rl "github.com/gen2brain/raylib-go/raylib"

func GetZombieSpawnerPositions(levelData *LevelColliderData) []rl.Vector2 {
	for _, layer := range levelData.ColliderLayers {
		if layer.Name == "Spawner_Zombie" {
			locations := getZombieSpawnerPostionsFromLayer(layer)

			return locations
		}
	}

	return make([]rl.Vector2, 0)
}

func getZombieSpawnerPostionsFromLayer(layer *ColliderLayer) []rl.Vector2 {
	var locations []rl.Vector2

	for x, cols := range layer.Data {
		for y, isSpawnerTile := range cols {
			if isSpawnerTile {
				location := rl.NewVector2(
					float32(x*WorldLevelData.TileWidth),
					float32(y*WorldLevelData.TileHeight),
				)

				locations = append(locations, location)
			}
		}
	}

	return locations
}
