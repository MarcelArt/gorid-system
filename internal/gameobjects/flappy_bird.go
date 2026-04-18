package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const jumpForce = 400

type FlappyBird struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Sprite   Spritesheet
}

func (g *FlappyBird) Start() {

}

func (g *FlappyBird) Update() {
	if rl.IsKeyPressed(rl.KeySpace) {
		g.Velocity.Y = -jumpForce
	}

	// rl.DrawRectangle(int32(g.Position.X), int32(g.Position.Y), 100, 100, rl.Blue)
	g.Sprite.DrawTile(0, g.Position)
}

func (g FlappyBird) GetPosition() rl.Vector2 {
	return g.Position
}

func (g FlappyBird) GetVelocity() rl.Vector2 {
	return g.Velocity
}

func (g *FlappyBird) SetPosition(pos rl.Vector2) {
	g.Position = pos
}

func (g *FlappyBird) SetVelocity(vel rl.Vector2) {
	g.Velocity = vel
}

func (g FlappyBird) GetSize() rl.Vector2 {
	return rl.Vector2{X: g.Sprite.TileSize.X * g.Sprite.Scale, Y: g.Sprite.TileSize.Y * g.Sprite.Scale}
}

func (g *FlappyBird) OnCollision(other Collidable) {
	// Handle collision - e.g., game over, reset, etc.
	// For now, just print a message
	println("Bird collided with obstacle!")
}

var (
	_ infrastructure.IGameObject = &FlappyBird{}
	_ Fallable                   = &FlappyBird{}
	_ Collidable                 = &FlappyBird{}
)
