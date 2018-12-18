package sc

import (
	"errors"
	"image"
	"math"
)

func CarvingX(img image.Image, f func(image.Image) [][]float64, debug bool) image.Image {
	fImg := flipImg(img)
	newFImg := CarvingY(fImg, f, debug)
	return flipImg(newFImg)
}

func CarvingY(img image.Image, f func(image.Image) [][]float64, debug bool) image.Image {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	energy := f(img)

	dir := []int{-1, 0, 1}

	dp := make([][]float64, dx)
	pre := make([][]int, dx)
	mark := make([][]int, dx)
	for i := 0; i < dx; i++ {
		dp[i] = make([]float64, dy)
		pre[i] = make([]int, dy)
		mark[i] = make([]int, dy)
		if i == 0 {
			continue
		}
		for j := 0; j < dy; j++ {
			tmp := math.Inf(1)
			last := j
			for _, d := range dir {
				lj := j + d
				if lj <= 0 || lj >= dy-1 {
					continue
				}
				tmp2 := dp[i-1][lj] + energy[i][j]
				if tmp2 < tmp {
					tmp = tmp2
					last = lj
				}
			}
			dp[i][j] = tmp
			pre[i][j] = last
		}
	}

	mnPos := -1
	mnVal := math.Inf(1)
	for i := 0; i < dy; i++ {
		if dp[dx-1][i] < mnVal {
			mnVal = dp[dx-1][i]
			mnPos = i
		}
	}
	//println(mnPos, mnVal)
	pos := mnPos
	posl := make([]int, dx)
	for i := dx - 1; i >= 0; i-- {
		mark[i][pos] = 1
		posl[i] = pos
		pos = pre[i][pos]
	}
	if debug {
		SavePNG(Overlay(img, mark), "seam.png")
	}

	newBounds := image.Rectangle{image.Pt(0, 0), image.Pt(dx, dy-1)}
	newImg := image.NewRGBA(newBounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			if j == posl[i] {
				continue
			} else if j < posl[i] {
				newImg.Set(i, j, img.At(i, j))
			} else {
				newImg.Set(i, j-1, img.At(i, j))
			}
		}
	}

	return newImg
}

func Resize(img image.Image, energyFunc func(image.Image) [][]float64, ndx, ndy int) image.Image {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	if ndx <= dx && ndy <= dy {
		if ndx == dx && ndy == dy {
			return img
		} else if ndx == dx {
			diff := dy - ndy
			for i := 0; i < diff; i++ {
				img = CarvingY(img, energyFunc, false)
			}
			return img
		} else if ndy == dy {
			diff := dx - ndx
			for i := 0; i < diff; i++ {
				img = CarvingX(img, energyFunc, false)
			}
			return img
		} else {
			panic(errors.New("cannot calc carving order"))
			return img
		}
	} else {
		panic(errors.New("cannot enlarge yet"))
		return img
	}
}
