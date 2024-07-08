package scenes

import (
	"github.com/Ollie-Ave/Zombie-Shooter/internal/entities"
	"github.com/Ollie-Ave/Zombie-Shooter/internal/levels"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameSceneHandler struct {
	levelData    *levels.LevelData
	colliderData *levels.LevelColliderData

	player    *entities.Player
	playerGun *entities.PlayerGun
	zombies   []*entities.Zombie
}

func NewGameSceneHandler() (*GameSceneHandler, error) {
	levelData, err := levels.LoadLevelData("test_level.json")

	if err != nil {
		return nil, err
	}

	colliderData := levels.LoadWorldColliderData(levels.WorldLevelData)

	playerStartingPos := rl.NewVector2(200, 350)

	player := entities.NewPlayer(playerStartingPos)
	playerGun := entities.NewPlayerGun(entities.PlayerEntity, rl.NewVector2(0, 0))

	zombies := []*entities.Zombie{
		entities.NewZombie(rl.NewVector2(100, 100)),
	}

	return &GameSceneHandler{
		levelData:    levelData,
		colliderData: colliderData,
		player:       player,
		playerGun:    playerGun,
		zombies:      zombies,
	}, nil
}

func (l *GameSceneHandler) Update() {
	levels.RenderLevelData(l.levelData, l.colliderData)

	l.player.Update()
	l.player.Render()
	l.playerGun.Update()
	l.playerGun.Render()

	for _, zombie := range l.zombies {
		zombie.Update()
		zombie.Render()
	}
}
