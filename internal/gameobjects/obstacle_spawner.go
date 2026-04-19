package gameobjects

import (
	"math"
	"math/rand"

	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObstacleSpawner struct {
	SpawnRate       float32 // in seconds
	Scene           infrastructure.MutableGameObjects
	timer           float32
	Obstacle        func(freeZone float32, freeZoneSize float32) *Obstacle
	CollisionSystem *CollisionSystem
}

func (g *ObstacleSpawner) Start() {
}

func (g *ObstacleSpawner) Update() {
	screenHeight := float64(rl.GetScreenHeight())
	rng := rand.Float64() * screenHeight
	rng = math.Max(150, float64(rng))
	rng = math.Min(rng, screenHeight-150)

	g.timer += rl.GetFrameTime()
	if g.timer > g.SpawnRate {
		spawnedObstacle := g.Obstacle(float32(rng), 150)
		obstacleCollider := &Collider{
			Object: spawnedObstacle,
			Rect: rl.NewRectangle(
				spawnedObstacle.Position.X,
				0,
				float32(spawnedObstacle.Sprite.TileSize.X*spawnedObstacle.Sprite.Scale),
				float32(screenHeight),
			),
		}
		g.CollisionSystem.AddCollider(obstacleCollider)
		g.Scene.AddGameObject(spawnedObstacle)
		g.Scene.AddGameObject(obstacleCollider)
		g.timer = 0
	}
}

var _ infrastructure.IGameObject = &ObstacleSpawner{}
