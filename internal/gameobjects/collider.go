package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collidable interface {
	infrastructure.IGameObject
	GetPosition() rl.Vector2
	GetSize() rl.Vector2
	OnCollision(other Collidable)
}

type Collider struct {
	Object   Collidable
	IsActive bool
}

func (c *Collider) Start() {}

func (c *Collider) Update() {
	if !c.IsActive {
		c.Object.Update()
	}
}

var _ infrastructure.IGameObject = &Collider{}
