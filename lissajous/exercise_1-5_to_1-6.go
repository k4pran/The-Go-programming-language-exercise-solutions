package lissajous

/*
 * EXERCISE 1.5
 * Change the Lissajous program's color palette to green on black, for added authenticity.
 * To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB, 0xFF}, where each pair
 * of hexadecimal digits represents the intensity of the red, green or blue component
 * of the pixel.

 * EXERCISE 1.6
 * Modify the Lissajous program to produce images in multiple colors by adding more values
 * to palette and then displaying them by changing the third argument of SetColorIndex in
 * some interesting way.
*/


import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"bufio"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x94, 0x00, 0xD3, 0xFF}, // Violet
	color.RGBA{0x4B, 0x00, 0x82, 0x00}, // Indigo
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Blue
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Green
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // Yellow
	color.RGBA{0xFF, 0x7F, 0x00, 0xFF}, // Orange
	color.RGBA{0x80, 0x00, 0x00, 0xFF}} // Red

const (
	blackIndex = 7
	greenIndex = 4
)

func main() {
	f, err := os.Create("lissajous.gif")
	if err == nil {
		out := bufio.NewWriter(f)
		lissajous(out)
	} else {
		log.Fatal(err)
	}

	f2, err := os.Create("lissajous_rainbow.gif")
	if err == nil {
		out := bufio.NewWriter(f2)
		lissajousRainbow(out)
	} else {
		log.Fatal(err)
	}
}

func lissajous(out io.Writer) {
	const (
		cycles= 5
		res= 0.001
		size= 100
		nframes= 64
		delay= 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func lissajousRainbow(out io.Writer) {
	const (
		cycles= 5
		res= 0.001
		size= 100
		nframes= 64
		delay= 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// Get random color from the rainbow.
			// Subtracting 1 from length and adding 1 afterwards avoids black.
			colorIndex := rand.Int() % (len(palette) - 1) + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

/*
====================================================================
						EXAMPLE OUTPUT
====================================================================

Find gif for exercise 1.5 at lissajous/lissajous.gif
Find gif for exercise 1.6 at lissajous/lissajous_rainbow.gif
 */