package gameobjects

import (
	"github.com/MarcelArt/gorid-system/internal/infrastructure"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Rect    rl.Rectangle
	Text    string
	OnClick func()
}

func (g *Button) Start() {

}

func (g *Button) Update() {
	mouse := rl.GetMousePosition()
	hovered := rl.CheckCollisionPointRec(mouse, g.Rect)

	color := rl.DarkGray
	if hovered {
		color = rl.Gray
	}

	rl.DrawRectangleRec(g.Rect, color)

	textWidth := rl.MeasureText(g.Text, 20)
	textX := int32(g.Rect.X + (g.Rect.Width-float32(textWidth))/2)
	textY := int32(g.Rect.Y + (g.Rect.Height / 2) - 10)

	rl.DrawText(g.Text, textX, textY, 20, rl.White)

	if hovered && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		g.OnClick()
	}
}

var _ infrastructure.IGameObject = &Button{}
