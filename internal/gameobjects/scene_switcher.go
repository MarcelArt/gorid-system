package gameobjects

import (
	"log"

	"github.com/MarcelArt/gorid-system/internal/game"
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneSwitcher struct {
	NextScene infrastructure.SceneConstructor
}

func NewSceneSwitcher(nextScene infrastructure.SceneConstructor) *SceneSwitcher {
	return &SceneSwitcher{
		NextScene: nextScene,
	}
}

func (g *SceneSwitcher) Start() {
}

func (g *SceneSwitcher) Update() {
	log.Println(1)
	if rl.IsKeyPressed(rl.KeySpace) {
		log.Println(2)
		game.Instance.LoadScene(g.NextScene)
	}
}
