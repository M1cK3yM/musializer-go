package audio

/*
#cgo LDFLAGS: -lraylib
#include <raylib.h>
#include <string.h>

extern void goAudioProcessor(void *buffer, unsigned int frames);
*/
import "C"

import (
	"math"
	"math/cmplx"
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Frame struct {
	Left  float32
	Right float32
}

const N uint = 256

var (
	in      []float32    = make([]float32, N)
	Out     []complex128 = make([]complex128, N)
	Max_amp float32
)

//export goAudioProcessor
func goAudioProcessor(buffer unsafe.Pointer, frames C.uint) {
	// src := (*[4800]Frame)(buffer)[:frames:frames]
	// capacity := (uint)(len(GlobalFrames))
	// fmt.Print("frames: ", frames, " capacity: ", capacity, " count: ", GlobalFramesCount, " len: ", len(src), "\n")
	// if (uint)(frames) <= capacity-GlobalFramesCount {
	// 	copy(GlobalFrames, src)
	// 	GlobalFramesCount += (uint)(frames)
	// } else if (uint)(frames) <= capacity {
	// 	copy(GlobalFrames[:capacity-(uint)(frames)], GlobalFrames[:(uint)(frames)])
	// 	copy(GlobalFrames[capacity-(uint)(frames):], src)
	// } else {
	// 	copy(GlobalFrames, src)
	// 	GlobalFramesCount = (uint)(frames)
	// }

	////////////////////////////////
	// if frames > (C.uint)(len(GlobalFrames)) {
	// 	frames = (C.uint)(len(GlobalFrames))
	// }
	//
	// copy(GlobalFrames, src)
	// GlobalFramesCount = (uint)(frames)

	if (uint)(frames) < N {
		return
	}
	fs := (*[N]Frame)(buffer)[:N:N]

	for i := range N {
		in[i] = (float32)((fs)[i].Left)
	}

	ditfft2(in, Out, (int)(N), 1)

	Max_amp = 0.0
	for i := range int(N) {
		a := Amp(Out[i])
		if Max_amp < a {
			Max_amp = a
		}
	}
}

func Amp(z complex128) float32 {
	a := math.Abs(real(z))
	b := math.Abs(imag(z))
	if a > b {
		return float32(a)
	}
	return float32(b)
}

func ditfft2(x []float32, y []complex128, n, s int) {
	if n == 1 {
		y[0] = complex(float64(x[0]), 0)
		return
	}
	ditfft2(x, y, n/2, 2*s)
	ditfft2(x[s:], y[n/2:], n/2, 2*s)
	for k := 0; k < n/2; k++ {
		tf := cmplx.Rect(1, -2*math.Pi*float64(k)/float64(n)) * y[k+n/2]
		y[k], y[k+n/2] = y[k]+tf, y[k]-tf
	}
}

func AttachAudioStreamProcessor(stream rl.AudioStream) {
	cStream := *(*C.AudioStream)(unsafe.Pointer(&stream))
	C.AttachAudioStreamProcessor(cStream, C.AudioCallback(C.goAudioProcessor))
}
