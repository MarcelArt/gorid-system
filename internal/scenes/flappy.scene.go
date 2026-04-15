package scenes

import (
	"github.com/MarcelArt/gorid-system/internal/gameobjects"
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FlappyScene struct {
	GameObjects []infrastructure.IGameObject
}

func NewFlappyScene() infrastructure.IScene {
	birdSpritesheet := gameobjects.Spritesheet{
		Texture: rl.LoadTexture("assets/Bird1-7.png"),
		TileSize: rl.Vector2{
			X: 16,
			Y: 16,
		},
		Scale: 3,
	}
	rl.SetTextureFilter(birdSpritesheet.Texture, rl.FilterPoint)

	flappyBird := &gameobjects.FlappyBird{
		Position: rl.Vector2{
			X: 350,
			Y: 100,
		},
		Sprite: birdSpritesheet,
	}
	flappyBirdGravity := &gameobjects.PhysicObject{
		Object: flappyBird,
		Gravity: rl.Vector2{
			X: 0,
			Y: 1,
		},
		GravMult: 800,
	}

	return &FlappyScene{
		GameObjects: []infrastructure.IGameObject{
			flappyBirdGravity,
			flappyBird,
		},
	}
}

func (s *FlappyScene) Update() {
	infrastructure.BaseUpdate(s, func() {
		rl.ClearBackground(rl.RayWhite)
	})
}

func (s *FlappyScene) GetGameObjects() []infrastructure.IGameObject {
	return s.GameObjects
}
