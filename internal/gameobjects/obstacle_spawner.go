package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObstacleSpawner struct {
	SpawnRate float32
	Scene     infrastructure.MutableGameObjects
	Timer     float32
}

func (g *ObstacleSpawner) Start() {
}

func (g *ObstacleSpawner) Update() {
	g.Timer += rl.GetFrameTime()
	if g.Timer > g.SpawnRate {
		g.Scene.AddGameObject(&Obstacle{})
		g.Timer = 0
	}
}

var _ infrastructure.IGameObject = &ObstacleSpawner{}
