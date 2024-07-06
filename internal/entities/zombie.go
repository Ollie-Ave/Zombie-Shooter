package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ZombieWidth     = 16
	ZombieHeight    = 16
	ZombieMaxHealth = 10
)

type Zombie struct {
	position rl.Vector2

	alive  bool
	health int
}

func NewZombie(position rl.Vector2) *Zombie {
	return &Zombie{
		position: position,
		health:   ZombieMaxHealth,
		alive:    true,
	}
}

func (z *Zombie) Update() {
	if z.health <= 0 {
		z.alive = false
	}
}

func (z *Zombie) Render() {
	if z.alive {
		rl.DrawRectangle(int32(z.position.X), int32(z.position.Y), ZombieWidth, ZombieHeight, rl.Green)
	}
}

func (z *Zombie) TakeDamage(damage int) {
	z.health -= damage
}

func (z *Zombie) GetHitbox() rl.Rectangle {
	return rl.NewRectangle(
		z.position.X,
		z.position.Y,
		ZombieWidth,
		ZombieHeight,
	)
}
