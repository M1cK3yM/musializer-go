package audio

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Music struct {
	music      rl.Music
	path       string
	wave       rl.Wave
	play       bool
	timePlayed float32
}

var m Music

func GetTimePlayed() float32 {
	return m.timePlayed
}

func GetWaveData() rl.Wave {
	return m.wave
}

func GetMusicStream() rl.AudioStream {
	return m.music.Stream
}

func InitMusic(path string) {
	m.music = rl.LoadMusicStream(path)
	m.path = path
	m.wave = rl.LoadWave(path)
	rl.PlayMusicStream(m.music)
	m.play = true
}

func UpdateMusicStream() {
	rl.UpdateMusicStream(m.music)
	m.timePlayed = rl.GetMusicTimePlayed(m.music) / rl.GetMusicTimeLength(m.music)
}

func TogglePlay() {
	m.play = !m.play
	if m.play {
		rl.PlayMusicStream(m.music)
	} else {
		rl.PauseMusicStream(m.music)
	}
}
