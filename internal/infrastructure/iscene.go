package infrastructure

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type IScene interface {
	Update()
	GetGameObjects() []IGameObject
}

type MutableGameObjects interface {
	AddGameObject(gameObject IGameObject)
	RemoveGameObject(gameObject IGameObject)
}

func BaseUpdate(s IScene, render func()) {
	rl.BeginDrawing()

	render()

	for _, gameObject := range s.GetGameObjects() {
		gameObject.Update()
	}

	rl.EndDrawing()
}

type SceneConstructor func() IScene
