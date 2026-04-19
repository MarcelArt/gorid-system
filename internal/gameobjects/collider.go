package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collidable interface {
	infrastructure.IGameObject
	OnCollision(other Collidable)
	GetPosition() rl.Vector2
}

type Collider struct {
	Object Collidable
	Rect   rl.Rectangle
}

func (g *Collider) Start() {
}

func (g *Collider) Update() {
	g.Rect.X = g.Object.GetPosition().X
	g.Rect.Y = g.Object.GetPosition().Y

}
