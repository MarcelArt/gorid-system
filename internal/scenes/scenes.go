package scenes

import "github.com/MarcelArt/gorid-system/internal/infrastructure"

// Implementation checker
var (
	_ infrastructure.IScene = NewLevelScene()
	_ infrastructure.IScene = NewMenuScene()
)
