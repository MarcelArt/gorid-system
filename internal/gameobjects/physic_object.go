package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Fallable interface {
	infrastructure.IGameObject
	GetPosition() rl.Vector2
	GetVelocity() rl.Vector2
	SetPosition(rl.Vector2)
	SetVelocity(rl.Vector2)
}

type PhysicObject struct {
	Gravity  rl.Vector2
	GravMult float32
	Object   Fallable
}

func (g *PhysicObject) Start() {
	// g.Object.GetXPosition()
	// g.Object.GetYPosition()
}

func (g *PhysicObject) Update() {
	dt := rl.GetFrameTime()

	pos := g.Object.GetPosition()
	vel := g.Object.GetVelocity()

	vel.X += g.GravMult * g.Gravity.X * dt
	pos.X += vel.X * dt

	vel.Y += g.GravMult * g.Gravity.Y * dt
	pos.Y += vel.Y * dt

	g.Object.SetPosition(pos)
	g.Object.SetVelocity(vel)
}
