package main

import (
	"fmt"
	"os"

	"github.com/Ollie-Ave/Zombie-Shooter/internal/entities"
	"github.com/Ollie-Ave/Zombie-Shooter/internal/levels"
	"github.com/Ollie-Ave/Zombie-Shooter/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowTargetFPS = 60
	WindowTitle     = "Zombie Shooter"

	WindowExitKey = rl.KeyCapsLock
)

var (
	WindowBackgroundColor = rl.LightGray
	Round                 = 0
)

func main() {
	setupWindow()
	setupWorld()

	for !rl.WindowShouldClose() {
		update()
	}

	rl.CloseWindow()
}

func setupWorld() {
	var err error

	levels.WorldLevelData, err = levels.LoadLevelData("test_level.json")

	if err != nil {
		panic(err)
	}

	levels.WorldColliderData = levels.LoadWorldColliderData(levels.WorldLevelData)

	startingPos := rl.NewVector2(200, 350)

	entities.PlayerEntity = entities.NewPlayer(startingPos)
	entities.PlayerGunEntity = entities.NewPlayerGun(entities.PlayerEntity, rl.NewVector2(0, 0))

	levelWidth := levels.WorldLevelData.TileWidth * levels.WorldLevelData.Width
	entities.CameraHandlerEntity = entities.NewCameraHandler(entities.PlayerEntity, levelWidth)

	spawnZombies()
}

func setupWindow() {
	rl.InitWindow(shared.WindowWidth, shared.WindowHeight, WindowTitle)

	rl.SetTargetFPS(WindowTargetFPS)
	rl.SetExitKey(WindowExitKey)

	rl.SetConfigFlags(rl.FlagVsyncHint)
}

func update() {

	allZombiesDead := true

	for _, zombie := range entities.ZombieEntities {
		if zombie.IsAlive() {
			allZombiesDead = false
			break
		}
	}

	if allZombiesDead {
		spawnZombies()
	}

	rl.BeginDrawing()

	rl.BeginMode2D(entities.CameraHandlerEntity.Camera)
	entities.CameraHandlerEntity.Update()

	rl.ClearBackground(WindowBackgroundColor)

	levels.RenderLevelData(levels.WorldLevelData, levels.WorldColliderData)

	entities.PlayerEntity.Update()
	entities.PlayerEntity.Render()
	entities.PlayerGunEntity.Update()
	entities.PlayerGunEntity.Render()

	for _, zombie := range entities.ZombieEntities {
		zombie.Update()
		zombie.Render()
	}

	rl.EndMode2D()

	renderUI()

	if shared.IsDebugMode() {
		renderFPS()
	}

	rl.EndDrawing()

	handleDebugMode()
}

func renderUI() {
	fpsText := fmt.Sprintf("Round: %d", Round)

	rl.DrawText(fpsText, 10, 10, 20, rl.White)
}

func renderFPS() {
	fpsText := fmt.Sprintf("FPS: %d", rl.GetFPS())

	rl.DrawText(fpsText, 10, 30, 20, rl.White)
}

func handleDebugMode() {
	if rl.IsKeyReleased(rl.KeyF3) {
		newDebugState := "true"

		if shared.IsDebugMode() {
			newDebugState = "false"
		}

		os.Setenv(shared.DebugModeEnvironmentVariable, newDebugState)
	}
}

func spawnZombies() {
	Round++

	zombiePositions := levels.GetZombieSpawnerPositions(levels.WorldColliderData)
	entities.ZombieEntities = make([]*entities.Zombie, len(zombiePositions))

	for index, position := range zombiePositions {
		zombieHealth := 1 + Round/3

		entities.ZombieEntities[index] = entities.NewZombie(position, zombieHealth)
	}
}
