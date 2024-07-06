package main

import (
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WindowWidth     = 1280
	WindowHeight    = 720
	WindowTargetFPS = 60
	WindowTitle     = "Project Kat"

	WindowExitKey = rl.KeyCapsLock

	DebugModeEnvironmentVariable = "DEBUG"
)

var (
	WindowBackgroundColor = rl.LightGray
)

func main() {
	fmt.Println("Hello World!")

	setupWindow()

	for !rl.WindowShouldClose() {
		update()
	}

	rl.CloseWindow()
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

	if os.Getenv(DebugModeEnvironmentVariable) == "true" {
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

		if os.Getenv(DebugModeEnvironmentVariable) == "true" {
			newDebugState = "false"
		}

		os.Setenv(DebugModeEnvironmentVariable, newDebugState)
	}
}
