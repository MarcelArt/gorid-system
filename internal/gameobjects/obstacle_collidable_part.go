package gameobjects

import rl "github.com/gen2brain/raylib-go/raylib"

type ObstacleCollidablePart struct {
	Rect rl.Rectangle
}

func (g *ObstacleCollidablePart) Update() {
}

func (g *ObstacleCollidablePart) OnCollision(other Collidable) {
}

func (g *ObstacleCollidablePart) GetPosition() rl.Vector2 {
	return rl.Vector2{
		X: g.Rect.X,
		Y: g.Rect.Y,
	}
}
