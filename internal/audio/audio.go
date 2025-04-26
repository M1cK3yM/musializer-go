package audio

import rl "github.com/gen2brain/raylib-go/raylib"

type Music struct {
	music      rl.Music
	play       bool
	timePlayed float32
}

var m Music

func GetTimePlayed() float32 {
	return m.timePlayed
}

func InitMusic(path string) {
	m.music = rl.LoadMusicStream(path)
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

func GetWaveData() []float32 {
	samples := rl.GetMusicTimePlayed(m.music) // You'd expand to real wave sampling
	return []float32{float32(samples)}
}

func GetFFTData() []float32 {
	// Real FFT data generation would be here
	return []float32{}
}
