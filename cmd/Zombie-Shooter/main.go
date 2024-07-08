package main

import (
	"fmt"
	"os"

	"github.com/Ollie-Ave/Zombie-Shooter/internal/scenes"
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

	SceneHandler *scenes.GameSceneHandler
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
	sceneHandler, err := scenes.NewGameSceneHandler()

	if err != nil {
		panic(err)
	}

	SceneHandler = sceneHandler
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

	SceneHandler.Update()

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
