package entities

import (
	"github.com/Ollie-Ave/Zombie-Shooter/internal/levels"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PlayerWidth  = 16
	PlayerHeight = 16
	PlayerSpeed  = 200
)

type Player struct {
	Position rl.Vector2
	Velocity rl.Vector2
}

func NewPlayer(startingPos rl.Vector2) *Player {
	return &Player{
		Position: startingPos,
	}
}

func (p *Player) Update() {
	nextPosition := p.Position

	deltaTime := rl.GetFrameTime()

	if rl.IsKeyDown(rl.KeyS) {
		nextPosition.Y += PlayerSpeed * deltaTime

		nextHitboxPosition := p.getHitboxForPosition(nextPosition)
		if levels.HitboxCollidesWithWorld(nextHitboxPosition) {
			nextPosition.Y -= PlayerSpeed * deltaTime
		}
	}
	if rl.IsKeyDown(rl.KeyW) {
		nextPosition.Y -= PlayerSpeed * deltaTime

		nextHitboxPosition := p.getHitboxForPosition(nextPosition)
		if levels.HitboxCollidesWithWorld(nextHitboxPosition) {
			nextPosition.Y += PlayerSpeed * deltaTime
		}
	}
	if rl.IsKeyDown(rl.KeyD) {
		nextPosition.X += PlayerSpeed * deltaTime

		nextHitboxPosition := p.getHitboxForPosition(nextPosition)
		if levels.HitboxCollidesWithWorld(nextHitboxPosition) {
			nextPosition.X -= PlayerSpeed * deltaTime
		}
	}
	if rl.IsKeyDown(rl.KeyA) {
		nextPosition.X -= PlayerSpeed * deltaTime

		nextHitboxPosition := p.getHitboxForPosition(nextPosition)
		if levels.HitboxCollidesWithWorld(nextHitboxPosition) {
			nextPosition.X += PlayerSpeed * deltaTime
		}
	}

	p.Position = nextPosition
}

func (p *Player) Render() {
	rl.DrawRectangle(
		int32(p.Position.X),
		int32(p.Position.Y),
		PlayerWidth,
		PlayerHeight,
		rl.Red)
}

func (p *Player) getHitboxForPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X,
		position.Y,
		PlayerWidth,
		PlayerHeight,
	)
}
