package entities

import rl "github.com/gen2brain/raylib-go/raylib"

func getHitboxCenter(hitbox rl.Rectangle) rl.Vector2 {
	return rl.NewVector2(
		hitbox.X+hitbox.Width/2,
		hitbox.Y+hitbox.Height/2,
	)
}
