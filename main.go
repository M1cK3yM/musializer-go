package main

import (
	"path/filepath"
	"strings"

	"github.com/M1cK3yM/musializer-go/internal/audio"
	"github.com/M1cK3yM/musializer-go/internal/config"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(config.ScreenWidth, config.ScreenHeight, "Musializer")
	rl.InitAudioDevice()
	rl.SetTargetFPS(config.FPS)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			audio.TogglePlay()
		}

		rl.BeginDrawing()
		rl.ClearBackground(config.BackgroundColor)

		rl.DrawRectangle(
			(int32)(rl.GetScreenWidth()-(int)(float64(rl.GetScreenWidth())*0.8)),
			int32(rl.GetScreenHeight())-50,
			400, 12, rl.LightGray)

		rl.DrawRectangle(
			(int32)(rl.GetScreenWidth()-(int)(float64(rl.GetScreenWidth())*0.8)),
			int32(rl.GetScreenHeight())-50,
			(int32)(audio.GetTimePlayed()*400.0), 12, rl.Maroon)

		rl.DrawRectangleLines(
			(int32)(rl.GetScreenWidth()-(int)(float64(rl.GetScreenWidth())*0.8)),
			int32(rl.GetScreenHeight())-50,
			400, 12, rl.Gray)

		if rl.IsFileDropped() {
			paths := rl.LoadDroppedFiles()
			for _, path := range paths {
				if strings.ToLower(filepath.Ext(path)) == ".mp3" ||
					strings.ToLower(filepath.Ext(path)) == ".wav" ||
					strings.ToLower(filepath.Ext(path)) == ".flac" {
					audio.InitMusic(path)
				}
			}
		}

		audio.UpdateMusicStream()
		rl.EndDrawing()
	}
}
