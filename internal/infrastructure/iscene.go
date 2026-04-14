package infrastructure

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type IScene interface {
	Update()
	GetGameObjects() []IGameObject
}

func BaseUpdate(s IScene, render func()) {
	rl.BeginDrawing()

	render()

	for _, gameObject := range s.GetGameObjects() {
		go gameObject.Update()
	}

	rl.EndDrawing()
}

type SceneConstructor func() IScene
