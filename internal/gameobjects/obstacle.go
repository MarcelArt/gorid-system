package gameobjects

import (
	"log"

	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TopPipe = 0
	// TopMiddlePipe    = 4
	MiddlePipe = 8
	// BottomMiddlePipe = 12
	BottomPipe = 12
)

type Obstacle struct {
	Sprite       Spritesheet
	Position     rl.Vector2
	FreeZone     float32
	FreeZoneSize float32
	// TopPipeCollider    *ObstacleCollidablePart
	// BottomPipeCollider *ObstacleCollidablePart
}

func ObstaclePrefab(sprite Spritesheet) func(float32, float32) *Obstacle {
	return func(freeZone float32, freeZoneSize float32) *Obstacle {
		return &Obstacle{
			Sprite:       sprite,
			Position:     rl.Vector2{X: 900, Y: 0},
			FreeZone:     freeZone,
			FreeZoneSize: freeZoneSize,
		}
	}
}

func (g *Obstacle) Start() {
}

func (g *Obstacle) Update() {
	dt := rl.GetFrameTime()
	g.Position.X -= 200 * dt

	// Draw bottom pipe (coming from floor up)
	g.drawFloorPipe()
	// Draw top pipe (coming from ceiling down)
	g.drawCeilingPipe()
}

func (g *Obstacle) GetPosition() rl.Vector2 {
	return g.Position
}

func (g *Obstacle) OnCollision(other Collidable) {
	log.Println("obstacle hit")
}

func (g *Obstacle) drawFloorPipe() {
	// Start from floor and build up
	screenHeight := rl.GetScreenHeight()
	currentY := float32(screenHeight)
	tileHeight := g.Sprite.TileSize.Y
	tileScale := g.Sprite.Scale

	// currentY -= tileHeight * tileScale

	// Middle segments until reaching the free zone (gap)
	for currentY > g.FreeZone+g.FreeZoneSize {
		g.Sprite.DrawTile(MiddlePipe+1, rl.Vector2{X: g.Position.X, Y: currentY})
		currentY -= tileHeight * tileScale
	}

	// Top cap
	g.Sprite.DrawTile(TopPipe+1, rl.Vector2{X: g.Position.X, Y: currentY})
}

func (g *Obstacle) drawCeilingPipe() {
	// Start from ceiling and build down to the gap
	tileHeight := g.Sprite.TileSize.Y * g.Sprite.Scale
	currentY := float32(0)

	// currentY += tileHeight

	// Middle segments from gap down (draw a reasonable number of segments)
	// screenHeight := rl.GetScreenHeight()
	for currentY < g.FreeZone-g.FreeZoneSize {
		g.Sprite.DrawTile(MiddlePipe, rl.Vector2{X: g.Position.X, Y: currentY})
		currentY += tileHeight
	}

	// Bottom cap
	g.Sprite.DrawTile(BottomPipe, rl.Vector2{X: g.Position.X, Y: currentY})
}

var _ infrastructure.IGameObject = &Obstacle{}
var _ Collidable = &Obstacle{}
