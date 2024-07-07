package entities

import (
	"github.com/Ollie-Ave/Zombie-Shooter/internal/levels"
	"github.com/Ollie-Ave/Zombie-Shooter/internal/shared"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ZombieWidth         = 16
	ZombieHeight        = 16
	ZombieMaxHealth     = 10
	ZombieMovementSpeed = 50
)

type Zombie struct {
	position               rl.Vector2
	lastSeenPlayerPosition rl.Vector2

	alive  bool
	health int
}

func NewZombie(position rl.Vector2) *Zombie {
	return &Zombie{
		position:               position,
		lastSeenPlayerPosition: position,
		health:                 ZombieMaxHealth,
		alive:                  true,
	}
}

func (z *Zombie) Update() {
	if z.health <= 0 {
		z.alive = false
	}

	if !z.alive {
		return
	}

	playerHitboxPos := getHitboxCenter(PlayerEntity.GetHitbox())
	zombieHitboxPos := getHitboxCenter(z.GetHitbox())

	var movementDirection rl.Vector2
	if z.canSeePlayer(zombieHitboxPos, playerHitboxPos) {
		movementDirection = rl.Vector2Normalize(rl.Vector2Subtract(playerHitboxPos, zombieHitboxPos))

		z.lastSeenPlayerPosition = playerHitboxPos
	} else {
		movementDirection = rl.Vector2Normalize(rl.Vector2Subtract(z.lastSeenPlayerPosition, zombieHitboxPos))
	}

	z.position = z.getNextPosition(movementDirection)
}

func (z *Zombie) canSeePlayer(zombieHitbox, playerHitbox rl.Vector2) bool {
	canSeePlayer := true

	for _, layerColliders := range levels.WorldColliderData.ColliderLayers {
		if layerColliders.IsCollidable {
			worldBlocksViewOfPlayer := worldBlocksViewOfPlayer(layerColliders, playerHitbox, zombieHitbox)

			canSeePlayer = !worldBlocksViewOfPlayer
		}
	}

	if canSeePlayer {
		if shared.IsDebugMode() {
			rl.DrawLine(
				int32(playerHitbox.X),
				int32(playerHitbox.Y),
				int32(zombieHitbox.X),
				int32(zombieHitbox.Y),
				rl.Red,
			)
		}
	}

	return canSeePlayer
}

func (z *Zombie) Render() {
	if !z.alive {
		return
	}

	rl.DrawRectangle(int32(z.position.X), int32(z.position.Y), ZombieWidth, ZombieHeight, rl.Green)
}

func (z *Zombie) TakeDamage(damage int) {
	z.health -= damage
}

func (z *Zombie) GetHitbox() rl.Rectangle {
	return z.getHitboxForPosition(z.position)
}

func worldBlocksViewOfPlayer(layer *levels.ColliderLayer, playerHitboxPosition, zombieHitboxPosition rl.Vector2) bool {
	for x, dataRow := range layer.Data {
		for y, collidable := range dataRow {
			if collidable {
				centerOfTile := rl.NewVector2(
					float32(x*levels.WorldLevelData.TileWidth+levels.WorldLevelData.TileWidth/2),
					float32(y*levels.WorldLevelData.TileWidth+levels.WorldLevelData.TileWidth/2),
				)

				if rl.CheckCollisionPointLine(
					centerOfTile,
					playerHitboxPosition,
					zombieHitboxPosition,
					10,
				) {
					return true
				}
			}
		}
	}

	return false
}

func (z *Zombie) getHitboxForPosition(position rl.Vector2) rl.Rectangle {
	return rl.NewRectangle(
		position.X,
		position.Y,
		ZombieWidth,
		ZombieHeight,
	)
}

func (z *Zombie) getNextPosition(direction rl.Vector2) rl.Vector2 {
	deltaTime := rl.GetFrameTime()

	nextPosition := z.position

	nextPositionYDiff := direction.Y * deltaTime * ZombieMovementSpeed
	nextPosition.Y += nextPositionYDiff
	nextHitboxPosition := z.getHitboxForPosition(nextPosition)

	if levels.HitboxCollidesWithWorld(nextHitboxPosition) {
		nextPosition.Y -= nextPositionYDiff
	}

	nextPositionXDiff := direction.X * deltaTime * ZombieMovementSpeed
	nextPosition.X += nextPositionXDiff
	nextHitboxPosition = z.getHitboxForPosition(nextPosition)

	if levels.HitboxCollidesWithWorld(nextHitboxPosition) {
		nextPosition.X -= nextPositionXDiff
	}

	return nextPosition
}
