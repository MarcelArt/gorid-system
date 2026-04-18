package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionSystem struct {
	colliders []*Collider
	Debug     bool
}

func (s *CollisionSystem) Start() {}

func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{
		colliders: make([]*Collider, 0),
	}
}

func (s *CollisionSystem) AddCollider(collider *Collider) {
	s.colliders = append(s.colliders, collider)
}

func (s *CollisionSystem) RemoveCollider(collider *Collider) {
	for i, c := range s.colliders {
		if c == collider {
			s.colliders = append(s.colliders[:i], s.colliders[i+1:]...)
			break
		}
	}
}

func (s *CollisionSystem) Update() {
	for i := 0; i < len(s.colliders); i++ {
		for j := i + 1; j < len(s.colliders); j++ {
			a := s.colliders[i]
			b := s.colliders[j]

			if a.IsActive && b.IsActive {
				if s.CheckCollision(a.Object, b.Object) {
					a.Object.OnCollision(b.Object)
					b.Object.OnCollision(a.Object)
				}
			}
		}
	}

	if s.Debug {
		s.DrawDebug()
	}
}

func (s *CollisionSystem) CheckCollision(a, b Collidable) bool {
	aPos := a.GetPosition()
	aSize := a.GetSize()
	bPos := b.GetPosition()
	bSize := b.GetSize()

	return rl.CheckCollisionRecs(
		rl.NewRectangle(aPos.X, aPos.Y, aSize.X, aSize.Y),
		rl.NewRectangle(bPos.X, bPos.Y, bSize.X, bSize.Y),
	)
}

func (s *CollisionSystem) DrawDebug() {
	for _, collider := range s.colliders {
		if !collider.IsActive {
			continue
		}

		pos := collider.Object.GetPosition()
		size := collider.Object.GetSize()

		rect := rl.NewRectangle(pos.X, pos.Y, size.X, size.Y)
		rl.DrawRectangleLines(int32(rect.X), int32(rect.Y), int32(rect.Width), int32(rect.Height), rl.Red)
	}
}

var _ infrastructure.IGameObject = &CollisionSystem{}
