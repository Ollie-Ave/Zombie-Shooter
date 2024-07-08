package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CameraZoom = 2.0

	cameraMovementOffset = 50
	maxCameraSpeed       = PlayerSpeed
	cameraSpeed          = 50
	cameraDrag           = 25
)

type CameraHandler struct {
	Camera     rl.Camera2D
	Player     *Player
	velocity   rl.Vector2
	levelWidth int
}

func NewCameraHandler(player *Player, levelWidth int) *CameraHandler {
	offset := rl.NewVector2(
		(float32(rl.GetScreenWidth())/2)-PlayerWidth,
		(float32(rl.GetScreenHeight())/2)-PlayerHeight,
	)

	camera := rl.NewCamera2D(
		offset,
		player.Position,
		0.0,
		CameraZoom,
	)

	return &CameraHandler{
		Camera:     camera,
		Player:     player,
		levelWidth: levelWidth,
		velocity:   rl.Vector2Zero(),
	}
}

func (c *CameraHandler) Update() {
	playerCenter := getHitboxCenter(c.Player.GetHitbox())
	screenCenter := rl.GetScreenToWorld2D(
		rl.NewVector2(
			float32(rl.GetScreenWidth()/2),
			float32(rl.GetScreenHeight()/2),
		),
		c.Camera,
	)

	if c.velocity.X > 0 {
		c.velocity.X = rl.Clamp(c.velocity.X-cameraDrag, 0, maxCameraSpeed)
	} else if c.velocity.X < 0 {
		c.velocity.X = rl.Clamp(c.velocity.X+cameraDrag, -maxCameraSpeed, 0)
	}

	if c.velocity.Y > 0 {
		c.velocity.Y = rl.Clamp(c.velocity.Y-cameraDrag, 0, maxCameraSpeed)
	} else if c.velocity.Y < 0 {
		c.velocity.Y = rl.Clamp(c.velocity.Y+cameraDrag, -maxCameraSpeed, 0)
	}

	if playerCenter.X > screenCenter.X+cameraMovementOffset {
		c.velocity.X = rl.Clamp(c.velocity.X+cameraSpeed, -maxCameraSpeed, maxCameraSpeed)
	} else if playerCenter.X < screenCenter.X-cameraMovementOffset {
		c.velocity.X = rl.Clamp(c.velocity.X-cameraSpeed, -maxCameraSpeed, maxCameraSpeed)
	}

	if playerCenter.Y > screenCenter.Y+cameraMovementOffset {
		c.velocity.Y = rl.Clamp(c.velocity.Y+cameraSpeed, -maxCameraSpeed, maxCameraSpeed)
	} else if playerCenter.Y < screenCenter.Y-cameraMovementOffset {
		c.velocity.Y = rl.Clamp(c.velocity.Y-cameraSpeed, -maxCameraSpeed, maxCameraSpeed)
	}

	deltaTime := rl.GetFrameTime()

	c.Camera.Target.X += c.velocity.X * deltaTime
	c.Camera.Target.X = float32(int(c.Camera.Target.X))

	c.Camera.Target.Y += c.velocity.Y * deltaTime
	c.Camera.Target.Y = float32(int(c.Camera.Target.Y))
}
