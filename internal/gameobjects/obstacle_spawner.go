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
		obstacle := g.Obstacle(float32(rng), 150)
		g.Scene.AddGameObject(obstacle)

		// Register obstacle's collider if CollisionSystem is available
		if g.CollisionSystem != nil {
			obstacleCollider := &Collider{
				Object:   obstacle,
				IsActive: true,
			}
			g.CollisionSystem.AddCollider(obstacleCollider)
		}

		g.timer = 0
	}
}

var _ infrastructure.IGameObject = &ObstacleSpawner{}
