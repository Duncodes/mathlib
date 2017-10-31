package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{100, 50, 0, 0}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	f, _ := os.Create("photo.png")
	e := png.Encode(f, m)
	//	_, e := f.Write(i)
	if e != nil {
		log.Println(e)
	}
}
