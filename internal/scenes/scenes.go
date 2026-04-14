package scenes

import "github.com/MarcelArt/gorid-system/internal/infrastructure"

var Map = map[string]func() infrastructure.IScene{
	"menu":  NewMenuScene,
	"level": NewLevelScene,
}

// Implementation checker
var (
	_ infrastructure.IScene = NewLevelScene()
	_ infrastructure.IScene = NewMenuScene()
)
