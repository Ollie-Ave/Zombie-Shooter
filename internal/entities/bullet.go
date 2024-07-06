package entities

import (
	"github.com/Ollie-Ave/Zombie-Shooter/internal/levels"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	BulletSize   = 5
	BulletSpeed  = 1000
	BulletDamage = 1
)

type Bullet struct {
	position  rl.Vector2
	direction rl.Vector2

	hasCollided bool
}

func NewBullet(initialPosition, direction rl.Vector2) *Bullet {
	return &Bullet{
		position:  initialPosition,
		direction: direction,
	}
}

func (b *Bullet) Render() {
	if !b.hasCollided {
		rec := rl.NewRectangle(
			b.position.X,
			b.position.Y,
			BulletSize,
			BulletSize,
		)

		origin := rl.NewVector2(0, 0)
		rotationAngle := calcuateAngleForDirectionVector(b.direction)

		rl.DrawRectanglePro(rec, origin, float32(rotationAngle), rl.Black)
	}
}

func (b *Bullet) Update() {
	if !b.hasCollided {
		b.position = b.getNextPosition()

		b.handleCollisions()
	}
}

func (b *Bullet) handleCollisions() {
	bulletHitbox := b.getHitbox()

	if levels.HitboxCollidesWithWorld(bulletHitbox) {
		b.hasCollided = true

		return
	}

	for _, zombie := range ZombieEntities {
		zombieHitbox := zombie.GetHitbox()

		if zombie.alive && rl.CheckCollisionRecs(bulletHitbox, zombieHitbox) {
			zombie.TakeDamage(BulletDamage)
			b.hasCollided = true

			return
		}
	}
}

func (b *Bullet) getNextPosition() rl.Vector2 {
	deltaTime := rl.GetFrameTime()

	return rl.NewVector2(
		b.position.X+BulletSpeed*-b.direction.X*deltaTime,
		b.position.Y+BulletSpeed*-b.direction.Y*deltaTime,
	)
}

func (b *Bullet) getHitbox() rl.Rectangle {
	return rl.NewRectangle(
		b.position.X,
		b.position.Y,
		BulletSize,
		BulletSize,
	)
}
