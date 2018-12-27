package SC

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

func LoadJPEG(s string) image.Image {
	f, err := os.Open(s)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	return img
}

func LoadPNG(s string) image.Image {
	f, err := os.Open(s)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	return img
}

func SavePNG(img image.Image, s string) {
	f, err := os.Create(s)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func SliceToImg(s [][]float64) image.Image {
	mx := float64(0)
	dx := len(s)
	dy := len(s[0])
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			if s[i][j] > mx {
				mx = s[i][j]
			}
		}
	}
	bounds := image.Rectangle{image.Point{0, 0}, image.Point{dx, dy}}
	newImg := image.NewGray(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			newImg.SetGray(i, j, color.Gray{uint8(s[i][j] / mx * 255)})
		}
	}
	return newImg
}

func IntSliceToImg(s [][]int) image.Image {
	mx := 0
	dx := len(s)
	dy := len(s[0])
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			if s[i][j] > mx {
				mx = s[i][j]
			}
		}
	}
	bounds := image.Rectangle{image.Point{0, 0}, image.Point{dx, dy}}
	newImg := image.NewGray(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			newImg.SetGray(i, j, color.Gray{uint8(s[i][j] * 255 / mx)})
		}
	}
	return newImg
}

func Overlay(img image.Image, s [][]int) image.Image {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	dx := bounds.Dx()
	dy := bounds.Dy()
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			if s[i][j] == 0 {
				newImg.Set(i, j, img.At(i, j))
			}
		}
	}
	return newImg
}

func flipImg(img image.Image) image.Image {
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	newBounds := image.Rectangle{image.Pt(0, 0), image.Pt(dy, dx)}
	newImg := image.NewRGBA(newBounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			newImg.Set(j, i, img.At(i, j))
		}
	}
	return newImg
}
