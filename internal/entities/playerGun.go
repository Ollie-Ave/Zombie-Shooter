package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	GunWidth  = 5
	GunHeight = 15
	GunOffset = -15
)

type PlayerGun struct {
	player *Player

	relativePosition rl.Vector2

	bullets []*Bullet
}

func NewPlayerGun(parent *Player, initialRelativePosition rl.Vector2) *PlayerGun {
	return &PlayerGun{
		player:           parent,
		relativePosition: initialRelativePosition,
	}
}

func (p *PlayerGun) Update() {
	p.updateGun()
	p.updateBullets()
}

func (p *PlayerGun) Render() {
	p.renderGun()
	p.renderBullets()
}

func (p *PlayerGun) getAbsolutePosition() rl.Vector2 {
	playerCenter := rl.Vector2Add(rl.NewVector2(PlayerWidth/2, PlayerHeight/2), p.player.Position)

	return rl.Vector2Add(playerCenter, p.relativePosition)
}

func (p *PlayerGun) updateBullets() {
	for _, bullet := range p.bullets {
		bullet.Update()
	}
}

func (p *PlayerGun) updateGun() {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		absolutePositon := p.getAbsolutePosition()
		mousePos := rl.GetMousePosition()

		mouseDirection := rl.Vector2Normalize(rl.Vector2Subtract(absolutePositon, mousePos))

		newBulletPosition := rl.Vector2Add(p.getAbsolutePosition(), rl.NewVector2(GunWidth/2, 0))

		bullet := NewBullet(newBulletPosition, mouseDirection)

		p.bullets = append(p.bullets, bullet)
	}
}

func (p *PlayerGun) renderBullets() {
	for _, bullet := range p.bullets {
		bullet.Render()
	}
}

func (p *PlayerGun) renderGun() {
	absolutePositon := p.getAbsolutePosition()

	mousePos := rl.GetMousePosition()

	mouseDirection := rl.Vector2Normalize(rl.Vector2Subtract(absolutePositon, mousePos))
	rotationAngle := calcuateAngleForDirectionVector(mouseDirection)

	gunOffsetVector := rl.NewVector2(mouseDirection.X*GunOffset, mouseDirection.Y*GunOffset)
	gunPosition := rl.Vector2Add(absolutePositon, gunOffsetVector)

	rec := rl.NewRectangle(
		gunPosition.X,
		gunPosition.Y,
		GunWidth,
		GunHeight,
	)

	origin := rl.NewVector2(GunWidth/2, 0)

	rl.DrawRectanglePro(rec, origin, float32(rotationAngle), rl.Brown)
}
