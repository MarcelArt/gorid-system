package scenes

import (
	"log"

	"github.com/MarcelArt/gorid-system/internal/gameobjects"
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LevelScene struct {
	GameObjects []infrastructure.IGameObject
}

func NewLevelScene() infrastructure.IScene {
	sceneSwitcher := gameobjects.NewSceneSwitcher(NewMenuScene)

	return &LevelScene{
		GameObjects: []infrastructure.IGameObject{
			sceneSwitcher,
		},
	}
}

func (s *LevelScene) Update() {
	infrastructure.BaseUpdate(s, func() {
		log.Println("level.go")
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
	})
}

func (s *LevelScene) GetGameObjects() []infrastructure.IGameObject {
	return s.GameObjects
}
