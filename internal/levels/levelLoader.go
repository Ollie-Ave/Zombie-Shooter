package levels

import (
	"encoding/json"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	assetsPath = os.Getenv("ASSETS_PATH")
)

func LoadLevelData(filePath string) (*LevelData, error) {
	var levelData *LevelData

	levelFilePath := fmt.Sprintf("%s/%s", assetsPath, filePath)
	levelFile, err := os.ReadFile(levelFilePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(levelFile, &levelData)

	if err != nil {
		return nil, err
	}

	for _, tileSet := range levelData.TileSets {
		tileTexurePath := fmt.Sprintf("%s/%s", assetsPath, tileSet.Image)

		_, err := os.Stat(tileTexurePath)

		if err != nil {
			return nil, err
		}

		tileSet.Texture = rl.LoadTexture(tileTexurePath)
	}

	return levelData, nil
}
