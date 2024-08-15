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
)

func main() {
	setupWindow()
	setupWorld()

	for !rl.WindowShouldClose() {
		update()
	}

	rl.CloseWindow()
}

func setupWindow() {
	rl.InitWindow(shared.WindowWidth, shared.WindowHeight, WindowTitle)

	rl.SetTargetFPS(WindowTargetFPS)
	rl.SetExitKey(WindowExitKey)

	rl.SetConfigFlags(rl.FlagVsyncHint)
}

func update() {

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
