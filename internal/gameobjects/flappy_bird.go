package gameobjects

import (
	"log"

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

func (g *FlappyBird) OnCollision(other Collidable) {
	// TODO: Handle collision
	log.Println("bird hit")
}

var (
	_ infrastructure.IGameObject = &FlappyBird{}
	_ Fallable                   = &FlappyBird{}
)
