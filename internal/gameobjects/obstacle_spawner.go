package gameobjects

import (
	"math"
	"math/rand"

	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObstacleSpawner struct {
	SpawnRate float32 // in seconds
	Scene     infrastructure.MutableGameObjects
	timer     float32
	Obstacle  func(freeZone float32, freeZoneSize float32) *Obstacle
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
		g.Scene.AddGameObject(g.Obstacle(float32(rng), 150))
		g.timer = 0
	}
}

var _ infrastructure.IGameObject = &ObstacleSpawner{}
