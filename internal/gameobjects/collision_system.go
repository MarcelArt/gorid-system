package gameobjects

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionSystem struct {
	Colliders []*Collider
	IsDebug   bool
}

func NewCollisionSystem(isDebug bool) *CollisionSystem {
	return &CollisionSystem{
		Colliders: make([]*Collider, 0),
		IsDebug:   isDebug,
	}
}

func (g *CollisionSystem) Start() {
}

func (g *CollisionSystem) Update() {
	g.checkCollisionsLoop()
	g.DrawDebug()
}

func (g *CollisionSystem) AddCollider(collider *Collider) {
	g.Colliders = append(g.Colliders, collider)
}

func (g *CollisionSystem) checkCollisionsLoop() {
	for _, collider := range g.Colliders {
		for _, other := range g.Colliders {
			if collider != other {
				if rl.CheckCollisionRecs(collider.Rect, other.Rect) {
					collider.Object.OnCollision(other.Object)
					other.Object.OnCollision(collider.Object)
				}
			}
		}
	}
}

func (g *CollisionSystem) DrawDebug() {
	if g.IsDebug {
		for _, collider := range g.Colliders {
			rl.DrawRectangleLines(int32(collider.Rect.X), int32(collider.Rect.Y), int32(collider.Rect.Width), int32(collider.Rect.Height), rl.Red)
		}
	}
}
