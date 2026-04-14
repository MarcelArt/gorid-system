package scenes

import (
	"log"

	"github.com/MarcelArt/gorid-system/internal/gameobjects"
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuScene struct {
	GameObjects []infrastructure.IGameObject
}

func NewMenuScene() infrastructure.IScene {
	// sceneSwitcher := gameobjects.NewSceneSwitcher(NewLevelScene)
	newGameBtn := &gameobjects.Button{
		Rect: rl.NewRectangle(300, 200, 200, 50),
		Text: "New Game",
		OnClick: func() {
			log.Println("New game clicked")
		},
	}
	settingsBtn := &gameobjects.Button{
		Rect: rl.NewRectangle(300, 270, 200, 50),
		Text: "Settings",
		OnClick: func() {
			log.Println("Settings clicked")
		},
	}
	exitBtn := &gameobjects.Button{
		Rect: rl.NewRectangle(300, 340, 200, 50),
		Text: "Exit",
		OnClick: func() {
			log.Println("Exit clicked")
		},
	}

	return &MenuScene{
		GameObjects: []infrastructure.IGameObject{
			newGameBtn,
			settingsBtn,
			exitBtn,
		},
	}
}

func (s *MenuScene) Update() {
	infrastructure.BaseUpdate(s, func() {
		log.Println("menu.go")
		rl.ClearBackground(rl.Black)
		rl.DrawText("Main Menu", 320, 100, 30, rl.White)
	})
}

func (s *MenuScene) GetGameObjects() []infrastructure.IGameObject {
	return s.GameObjects
}
