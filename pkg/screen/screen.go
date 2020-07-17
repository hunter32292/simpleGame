package screen

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// DrawScreen - Draw the game screen
func DrawScreen() {
	rl.InitWindow(800, 800, "Simple Game")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Welcome to the Simple Game!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

}

// DrawIntroScreen - Draw the game screen
func DrawIntroScreen() {
	rl.InitWindow(800, 800, "Simple Game")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Welcome to the Simple Game!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	time.Sleep(time.Second * 4)

	rl.CloseWindow()
}
