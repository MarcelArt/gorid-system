package gameobjects

import "github.com/MarcelArt/gorid-system/internal/infrastructure"

type ObstacleSpawner struct {
	SpawnRate int
}

func (g *ObstacleSpawner) Start() {
}

func (g *ObstacleSpawner) Update() {
}

var _ infrastructure.IGameObject = &ObstacleSpawner{}
