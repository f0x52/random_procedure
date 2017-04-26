package main

import (
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"math/rand"
	"os"
	"time"
)

const (
	sizeX   = 3360
	sizeY   = 1080
	lines   = 20
	length  = 4000
	random  = false
	delay   = 5
	animate = false
)

var Map [sizeX][sizeY]int

//var img = image.NewRGBA(image.Rect(0, 0, sizeX, sizeY))
var img = image.NewPaletted(image.Rect(0, 0, sizeX, sizeY), palette)
var save = 0

var x, y int

var palette = []color.Color{
	color.RGBA{21, 22, 25, 255},
	color.RGBA{230, 230, 230, 255},
	color.RGBA{107, 238, 255, 255},
}

var frames []*image.Paletted
var delays []int

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for x := 0; x < sizeX; x++ {
		var row [sizeY]int
		for y := 0; y < sizeY; y++ {
			row[y] = 0
			img.Set(x, y, color.RGBA{21, 22, 25, 255}) //bg
		}
		Map[x] = row
	}
	for a := 0; a < lines; a++ {
		x := r.Intn(sizeX)
		y := r.Intn(sizeY)
		action := r.Intn(7) + 1
		for i := 0; i < length; i++ {
			c := r.Intn(100)
			if c > 98 {
				if random {
					action = r.Intn(7) + 1
				} else {
					action = action + 1
					if action == 9 {
						action = 1
					}
				}
			}

			switch action {
			case 1:
				y--
			case 2:
				x++
				y--
			case 3:
				x++
			case 4:
				x++
				y++
			case 5:
				y++
			case 6:
				x--
				y++
			case 7:
				x--
			case 8:
				x--
				y--
			}

			if x < sizeX && y < sizeY && x >= 0 && y >= 0 {
				Map[x][y] = r.Intn(2) + 1
				draw(x, y, r.Intn(2)+1)
			}
		}
	}
	//draw()
	//file, _ := os.Create("stage/out" + time.Now().Format("150405") + ".png")
	//defer file.Close()
	//png.Encode(file, img)
	//save = 0
	if animate {
		file, _ := os.Create("out.gif")
		gif.EncodeAll(file, &gif.GIF{Image: frames, Delay: delays})
	} else {
		file, _ := os.Create("out.png")
		defer file.Close()
		png.Encode(file, img)
	}
}

func draw(x int, y int, col int) {
	//	for x := 0; x < sizeX; x++ {
	//		for y := 0; y < sizeY; y++ {
	switch col {
	//case 0:
	//	img.Set(x, y, color.RGBA{21, 22, 25, 255}) //bg
	case 1:
		img.Set(x, y, color.RGBA{230, 230, 230, 255})
	case 2:
		img.Set(x, y, color.RGBA{107, 238, 255, 255})
	}
	//}
	//}
	if animate {
		save++
		if save == 10 {
			frames = append(frames, img)
			delays = append(delays, delay)
			save = 0
			img = image.NewPaletted(image.Rect(0, 0, sizeX, sizeY), palette)
			for x := 0; x < sizeX; x++ {
				for y := 0; y < sizeY; y++ {
					switch Map[x][y] {
					case 1:
						img.Set(x, y, color.RGBA{230, 230, 230, 255})
					case 2:
						img.Set(x, y, color.RGBA{107, 238, 255, 255})
					}
				}
			}

		}
	}
}
