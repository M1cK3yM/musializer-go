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

	musicLoaded := false
	processorsAttached := false

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			audio.TogglePlay()
			processorsAttached = false
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
					musicLoaded = true
				}
			}
			processorsAttached = false
		}

		h := rl.GetRenderHeight()
		w := rl.GetRenderWidth()

		if musicLoaded {
			if !processorsAttached {
				audio.AttachAudioStreamProcessor(audio.GetMusicStream())
				processorsAttached = true
			}
			// cellWidth := (int)(w) / int(audio.GlobalFramesCount)
			// for idx, gl := range audio.GlobalFrames {
			// 	t := (float32)((float32(h/2) * (float32(gl.Left))))
			// 	if t > 0 {
			// 		rl.DrawRectangle(int32(idx*cellWidth), (int32)(h/2)-(int32)(t), int32(cellWidth), (int32)(t), rl.White)
			// 	} else {
			// 		rl.DrawRectangle(int32(idx*cellWidth), (int32)(h/2), int32(cellWidth), (int32)(t), rl.White)
			// 	}
			// }
			cellWidth := (int)(w) / int(audio.N)
			for idx := range audio.N {
				t := (float32)((float32)(audio.Amp(audio.Out[idx])) / (float32)(audio.Max_amp) * (float32(h / 2)))
				rl.DrawRectangle(int32(idx*(uint)(cellWidth)), (int32)(h/2)-(int32)(t), int32(cellWidth), (int32)(t), rl.White)
			}
			drawSpectrogram()
		}

		audio.UpdateMusicStream()
		rl.EndDrawing()
	}
}

func drawSpectrogram() {
}
