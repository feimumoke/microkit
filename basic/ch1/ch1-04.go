package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

/**
GIF动画
*/
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func Lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 400
		nframes = 64
		delay   = 8
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main41() {
	rand.Seed(time.Now().UTC().UnixNano())
	Lissajous(os.Stdout)
}

// ----------------------------
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		Lissajous(writer)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
