package main

import (
	"github.com/MarcelArt/gorid-system/internal/game"
	"github.com/MarcelArt/gorid-system/internal/scenes"
)

func main() {
	// scene := scenes.NewFlappyScene()
	scene := scenes.NewMenuScene()

	game.InitGame(scene)
	game.Instance.Start()
}

// type Button struct {
// 	Rect rl.Rectangle
// 	Text string
// }

// func (b Button) Draw() bool {
// 	mouse := rl.GetMousePosition()
// 	hovered := rl.CheckCollisionPointRec(mouse, b.Rect)

// 	// Change color on hover
// 	color := rl.DarkGray
// 	if hovered {
// 		color = rl.Gray
// 	}

// 	// Draw button
// 	rl.DrawRectangleRec(b.Rect, color)

// 	// Center text
// 	textWidth := rl.MeasureText(b.Text, 20)
// 	textX := int32(b.Rect.X + (b.Rect.Width-float32(textWidth))/2)
// 	textY := int32(b.Rect.Y + (b.Rect.Height / 2) - 10)

// 	rl.DrawText(b.Text, textX, textY, 20, rl.White)

// 	// Return true if clicked
// 	return hovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)
// }

// func main() {
// 	screenWidth := int32(800)
// 	screenHeight := int32(600)

// 	rl.InitWindow(screenWidth, screenHeight, "Main Menu")
// 	rl.SetWindowState(rl.FlagWindowResizable)
// 	defer rl.CloseWindow()

// 	rl.SetTargetFPS(60)

// 	// Create buttons
// 	buttons := []Button{
// 		{rl.NewRectangle(300, 200, 200, 50), "New Game"},
// 		{rl.NewRectangle(300, 270, 200, 50), "Settings"},
// 		{rl.NewRectangle(300, 340, 200, 50), "Exit"},
// 	}

// 	for !rl.WindowShouldClose() {
// 		rl.BeginDrawing()
// 		rl.ClearBackground(rl.Black)

// 		rl.DrawText("Main Menu", 320, 100, 30, rl.White)

// 		// Handle buttons
// 		for _, btn := range buttons {
// 			if btn.Draw() {
// 				switch btn.Text {
// 				case "New Game":
// 					println("Start Game")
// 				case "Settings":
// 					println("Open Settings")
// 				case "Exit":
// 					rl.CloseWindow()
// 				}
// 			}
// 		}

// 		rl.EndDrawing()
// 	}
// }
