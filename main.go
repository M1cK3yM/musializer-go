package main

import (
	"github.com/M1cK3yM/musializer-go/internal/audio"
	"github.com/M1cK3yM/musializer-go/internal/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Musializer")
	rl.InitAudioDevice()
	rl.SetTargetFPS(config.FPS)
	defer rl.CloseWindow()

	audio.InitMusic("/path/to/music")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(config.BackgroundColor)

		audio.UpdateMusicStream()

		rl.EndDrawing()
	}
}
