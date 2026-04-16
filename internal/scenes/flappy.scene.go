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
	flappyBird, flappyBirdGravity := setupFlappyPlayer()

	scene := &FlappyScene{
		GameObjects: []infrastructure.IGameObject{
			flappyBirdGravity,
			flappyBird,
		},
	}

	obstacle := setupObstacle(scene)
	scene.AddGameObject(obstacle)

	return scene
}

func (s *FlappyScene) Update() {
	infrastructure.BaseUpdate(s, func() {
		rl.ClearBackground(rl.RayWhite)
	})
}

func (s *FlappyScene) GetGameObjects() []infrastructure.IGameObject {
	return s.GameObjects
}

func (s *FlappyScene) AddGameObject(gameObject infrastructure.IGameObject) {
	s.GameObjects = append(s.GameObjects, gameObject)
}

func (s *FlappyScene) RemoveGameObject(gameObject infrastructure.IGameObject) {
	// TODO: implement
}

func setupFlappyPlayer() (*gameobjects.FlappyBird, *gameobjects.PhysicObject) {
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

	return flappyBird, flappyBirdGravity
}

func setupObstacle(s *FlappyScene) *gameobjects.ObstacleSpawner {
	obstacleSpritesheet := gameobjects.Spritesheet{
		Texture: rl.LoadTexture("assets/PipeStyle1.png"),
		TileSize: rl.Vector2{
			X: 32,
			Y: 20,
		},
		Scale: 3,
	}

	obstacleSpawner := gameobjects.ObstacleSpawner{
		SpawnRate: 2,
		Scene:     s,
		Obstacle:  gameobjects.ObstaclePrefab(obstacleSpritesheet),
	}

	return &obstacleSpawner

	// return &gameobjects.Obstacle{
	// 	Sprite: obstacleSpritesheet,
	// 	Position: rl.Vector2{
	// 		X: 550,
	// 		Y: 600,
	// 	},
	// 	FreeZone:     250,
	// 	FreeZoneSize: 150,
	// }
}
