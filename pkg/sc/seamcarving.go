package sc

import (
	"errors"
	"image"
	"math"
)

func CarvingY(img image.Image, f func(image.Image, int) [][]float64, debug bool) (image.Image, float64) {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	energy := f(img, 2)

	dir := []int{-1, 0, 1}

	dp := make([][]float64, dx)
	pre := make([][]int, dx)
	for i := 0; i < dx; i++ {
		dp[i] = make([]float64, dy)
		pre[i] = make([]int, dy)
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

	pos := mnPos
	posl := make([]int, dx)

	if debug {
		mark := make([][]int, dx)
		for i := 0; i < dx; i++ {
			mark[i] = make([]int, dy)
		}
		for i := dx - 1; i >= 0; i-- {
			mark[i][pos] = 1
			posl[i] = pos
			pos = pre[i][pos]
		}
		SavePNG(Overlay(img, mark), "seam_y.png")
	}

	for i := dx - 1; i >= 0; i-- {
		posl[i] = pos
		pos = pre[i][pos]
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

	return newImg, mnVal
}

func CarvingX(img image.Image, f func(image.Image, int) [][]float64, debug bool) (image.Image, float64) {
	bounds := img.Bounds()

	// transpose
	dx := bounds.Dy()
	dy := bounds.Dx()

	energy := f(img, 1)

	dir := []int{-1, 0, 1}

	dp := make([][]float64, dx)
	pre := make([][]int, dx)
	for i := 0; i < dx; i++ {
		dp[i] = make([]float64, dy)
		pre[i] = make([]int, dy)

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
				tmp2 := dp[i-1][lj] + energy[j][i]
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

	if debug {
		mark := make([][]int, dy)
		for i := 0; i < dy; i++ {
			mark[i] = make([]int, dx)
		}
		for i := dx - 1; i >= 0; i-- {
			mark[pos][i] = 1
			posl[i] = pos
			pos = pre[i][pos]
		}
		SavePNG(Overlay(img, mark), "seam_x.png")
	}

	for i := dx - 1; i >= 0; i-- {
		posl[i] = pos
		pos = pre[i][pos]
	}

	newBounds := image.Rectangle{image.Pt(0, 0), image.Pt(dy-1, dx)}
	newImg := image.NewRGBA(newBounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			if j == posl[i] {
				continue
			} else if j < posl[i] {
				newImg.Set(j, i, img.At(j, i))
			} else {
				newImg.Set(j-1, i, img.At(j, i))
			}
		}
	}

	return newImg, mnVal
}

func Resize(img image.Image, energyFunc func(image.Image, int) [][]float64, ndx, ndy int) image.Image {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	if ndx <= dx && ndy <= dy {
		if ndx == dx && ndy == dy {
			return img
		} else if ndx == dx {
			diff := dy - ndy
			bar := Pbar{}
			bar.Init(diff)
			for i := 0; i < diff; i++ {
				img, _ = CarvingY(img, energyFunc, false)
				bar.Step(i + 1)
			}
			return img
		} else if ndy == dy {
			fImg := flipImg(img)
			retImg := Resize(fImg, energyFunc, ndy, ndx)
			return flipImg(retImg)
		} else {
			diffx := dx - ndx + 1
			diffy := dy - ndy + 1

			if diffx < diffy {
				fImg := flipImg(img)
				retImg := Resize(fImg, energyFunc, ndy, ndx)
				return flipImg(retImg)
			} else {
				I := make([][]image.Image, 2)
				I[0] = make([]image.Image, diffy)
				I[1] = make([]image.Image, diffy)
				T := make([][]float64, 2)
				T[0] = make([]float64, diffy)
				T[1] = make([]float64, diffy)
				I[0][0] = img
				T[0][0] = 0
				bar := Pbar{}
				bar.Init(diffx * diffy)
				for i := 0; i < diffx; i++ {
					for j := 0; j < diffy; j++ {
						if i == 0 && j == 0 {
							continue
						}
						T[i&1][j] = math.Inf(1)
						if i > 0 {
							img, val := CarvingX(I[i&1^1][j], energyFunc, false)
							if val+T[i&1^1][j] < T[i&1][j] {
								T[i&1][j] = val + T[i&1^1][j]
								I[i&1][j] = img
							}
						}
						if j > 0 {
							img, val := CarvingY(I[i&1][j-1], energyFunc, false)
							if val+T[i&1][j-1] < T[i&1][j] {
								T[i&1][j] = val + T[i&1][j-1]
								I[i&1][j] = img
							}
						}
						bar.Step(i*diffy + j + 1)
					}
				}
				return I[diffx&1^1][diffy-1]
			}
		}
	} else {
		panic(errors.New("cannot enlarge yet"))
		return img
	}
}
