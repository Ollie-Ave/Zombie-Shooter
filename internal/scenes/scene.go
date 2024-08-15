package scenes

import (
	"github.com/Ollie-Ave/Zombie-Shooter/internal/entities"
	"github.com/Ollie-Ave/Zombie-Shooter/internal/levels"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneHandler interface {
	SetupScene() error
	UpdateBeforeDraw()
	UpdateDuringDraw()
}

type GameScene struct {
	Round int
}

func NewGameScene() *GameScene {
	return &GameScene{
		Round: 0,
	}
}

func (g *GameScene) SetupScene() error {
	var err error

	levels.WorldLevelData, err = levels.LoadLevelData("test_level.json")

	if err != nil {
		return err
	}

	levels.WorldColliderData = levels.LoadWorldColliderData(levels.WorldLevelData)

	startingPos := rl.NewVector2(200, 350)

	entities.PlayerEntity = entities.NewPlayer(startingPos)
	entities.PlayerGunEntity = entities.NewPlayerGun(entities.PlayerEntity, rl.NewVector2(0, 0))

	levelWidth := levels.WorldLevelData.TileWidth * levels.WorldLevelData.Width
	entities.CameraHandlerEntity = entities.NewCameraHandler(entities.PlayerEntity, levelWidth)

	g.spawnZombies()

	return nil
}

func (g *GameScene) UpdateBeforeDraw() {
	allZombiesDead := true

	for _, zombie := range entities.ZombieEntities {
		if zombie.IsAlive() {
			allZombiesDead = false
			break
		}
	}

	if allZombiesDead {
		g.spawnZombies()
	}
}

func (*GameScene) UpdateDuringDraw() {
}

func (g *GameScene) spawnZombies() {
	g.Round++

	zombiePositions := levels.GetZombieSpawnerPositions(levels.WorldColliderData)
	entities.ZombieEntities = make([]*entities.Zombie, len(zombiePositions))

	for index, position := range zombiePositions {
		zombieHealth := 1 + g.Round/3

		entities.ZombieEntities[index] = entities.NewZombie(position, zombieHealth)
	}
}
