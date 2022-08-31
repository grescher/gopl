// Modify the Lissajous program to produce images in multiple colors
// by adding more values to `palette` and then displaying them by changing
// the third argument of `SetColorIndex` in some interesting way.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x1b, 0x26, 0x31, 0xff}, // dark grey
	color.RGBA{0xf9, 0xe7, 0x9f, 0xff}, // yellow
	color.RGBA{0xfa, 0xd7, 0xa0, 0xff}, // orange
	color.RGBA{0xf5, 0xb7, 0xb1, 0xff}, // red
	color.RGBA{0xd2, 0xb4, 0xde, 0xff}, // violet
	color.RGBA{0xae, 0xd6, 0xf1, 0xff}, // blue
	color.RGBA{0xab, 0xeb, 0xc6, 0xff}, // green
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms unit
		colors  = 6     // number of drawing colors in palette
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8((i/8)%colors+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
