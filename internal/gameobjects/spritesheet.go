package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Spritesheet struct {
	Texture  rl.Texture2D
	TileSize rl.Vector2
	Scale    float32
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

	dest := rl.Rectangle{
		X:      pos.X,
		Y:      pos.Y,
		Width:  g.TileSize.X * g.Scale,
		Height: g.TileSize.Y * g.Scale,
	}

	origin := rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(g.Texture, source, dest, origin, 0, rl.White)
}

var _ infrastructure.IGameObject = &Spritesheet{}

// Builders
type SpritesheetBuilder struct {
	Spritesheet Spritesheet
}

func NewSpritesheetBuilder(path string) *SpritesheetBuilder {
	return &SpritesheetBuilder{
		Spritesheet: Spritesheet{
			Texture: rl.LoadTexture(path),
			Scale:   1,
		},
	}
}

func (b *SpritesheetBuilder) SetScale(scale float32) *SpritesheetBuilder {
	b.Spritesheet.Scale = scale
	return b
}

func (b *SpritesheetBuilder) SetTileSize(tileSize rl.Vector2) *SpritesheetBuilder {
	b.Spritesheet.TileSize = tileSize
	return b
}

func (b *SpritesheetBuilder) Build() *Spritesheet {
	return &b.Spritesheet
}
