package snippet

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}
const (
	whiteIndex = 0
	blackIndex = 1
)

func Lissajous(out io.Writer) {
	const (
		cycles = 5
		resolution = 0.001
		size = 100
		numberOfFrames = 64
		delay = 8
	)

	frequency := rand.Float64() * 3.0
	animation := gif.GIF{LoopCount: numberOfFrames}
	phase := 0.0

	for i := 0; i < numberOfFrames; i++ {
		rectangle := image.Rect(0, 0, 2*size+1, 2*size+1)
		image := image.NewPaletted(rectangle, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += resolution {
			x := math.Sin(t)
			y := math.Sin(t*frequency + phase)
			image.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, image)
	}
	gif.EncodeAll(out, &animation)
}