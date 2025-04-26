package audio

import rl "github.com/gen2brain/raylib-go/raylib"

var music rl.Music

func InitMusic(path string) {
	music = rl.LoadMusicStream(path)
	rl.PlayMusicStream(music)
}

func UpdateMusicStream() {
	rl.UpdateMusicStream(music)
}

func GetWaveData() []float32 {
	samples := rl.GetMusicTimePlayed(music) // You'd expand to real wave sampling
	return []float32{float32(samples)}
}

func GetFFTData() []float32 {
	// Real FFT data generation would be here
	return []float32{}
}
