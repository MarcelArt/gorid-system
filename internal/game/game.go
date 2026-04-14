package game

import (
	"log"

	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	ActiveScene infrastructure.IScene

	Width     int32
	Height    int32
	Title     string
	TargetFPS int32
}

var Instance *Game

func InitGame(scene infrastructure.IScene) {
	if Instance == nil {
		Instance = &Game{
			ActiveScene: scene,
			Width:       800,
			Height:      450,
			Title:       "raylib [core] example - basic window",
			TargetFPS:   60,
		}
	} else {
		Instance.ActiveScene = scene
	}
}

func (g *Game) Start() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.ActiveScene.Update()
	}
}

func (g *Game) UpdateScene(scene infrastructure.SceneConstructor) {
	log.Println(3)
	g.ActiveScene = scene()
}
