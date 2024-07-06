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
	WindowWidth     = 1280
	WindowHeight    = 720
	WindowTargetFPS = 60
	WindowTitle     = "Project Kat"

	WindowExitKey = rl.KeyCapsLock
)

var (
	WindowBackgroundColor = rl.LightGray
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

	entities.ZombieEntities = []*entities.Zombie{
		entities.NewZombie(rl.NewVector2(100, 100)),
	}
}

func setupWindow() {
	rl.InitWindow(WindowWidth, WindowHeight, WindowTitle)

	rl.SetTargetFPS(WindowTargetFPS)
	rl.SetExitKey(WindowExitKey)

	rl.SetConfigFlags(rl.FlagVsyncHint)
}

func update() {
	rl.BeginDrawing()

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

	if shared.IsDebugMode() {
		renderFPS()
	}

	rl.EndDrawing()

	handleDebugMode()
}

func renderFPS() {
	fpsText := fmt.Sprintf("FPS: %d", rl.GetFPS())

	rl.DrawText(fpsText, 10, 10, 20, rl.White)
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
