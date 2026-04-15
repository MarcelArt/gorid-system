package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Spritesheet struct {
	Texture  rl.Texture2D
	TileSize rl.Vector2
}

func (g *Spritesheet) Start() {
}

func (g *Spritesheet) Update() {
}

func (g *Spritesheet) DrawTile(tileIndex int, pos rl.Vector2) {
	tilesPerRow := g.Texture.Width / int32(g.TileSize.X)

	srcX := (tileIndex % int(tilesPerRow)) * int(g.TileSize.X)
	srcY := (tileIndex / int(tilesPerRow)) * int(g.TileSize.Y)

	source := rl.Rectangle{
		X:      float32(srcX),
		Y:      float32(srcY),
		Width:  g.TileSize.X,
		Height: g.TileSize.Y,
	}

	rl.DrawTextureRec(g.Texture, source, pos, rl.White)
}

var _ infrastructure.IGameObject = &Spritesheet{}
