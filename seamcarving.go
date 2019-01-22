package SC

import (
	"errors"
	"golang.org/x/image/colornames"
	"image"
	"math"
)

func setMask(energyp *[][]float64, imgMask image.Image, dx, dy int) {
	if imgMask != nil {
		energy := *energyp
		for i := 0; i < dx; i++ {
			for j := 0; j < dy; j++ {
				r, g, b, _ := imgMask.At(i, j).RGBA()
				if r >= 200 && g <= 20 && b <= 20 { // 红色 删去
					energy[i][j] = -1e10
				}
				if g >= 200 && r <= 20 && b <= 20 { // 绿色 留下
					energy[i][j] = 1e10
				}
			}
		}
	}
}

func carvingY(img image.Image, f func(image.Image, int) [][]float64, posp *[][]int, imgMask image.Image) (image.Image, *[][]int, float64, image.Image) {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	energy := f(img, 2)
	setMask(&energy, imgMask, dx, dy)

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

	for i := dx - 1; i >= 0; i-- {
		posl[i] = pos
		pos = pre[i][pos]
	}
	newBounds := image.Rectangle{image.Pt(0, 0), image.Pt(dx, dy-1)}
	newImg := image.NewRGBA(newBounds)

	retp := (*[][]int)(nil)

	if posp != nil {
		posa := *posp
		reta := make([][]int, dx)
		for i := 0; i < dx; i++ {
			reta[i] = make([]int, dy-1)
		}
		retp = &reta

		for i := 0; i < dx; i++ {
			for j := 0; j < dy; j++ {
				if j == posl[i] {
					continue
				} else if j < posl[i] {
					reta[i][j] = posa[i][j]
				} else {
					reta[i][j-1] = posa[i][j]
				}
			}
		}
	}

	if imgMask != nil {
		newMask := image.NewRGBA(newBounds)
		for i := 0; i < dx; i++ {
			for j := 0; j < dy; j++ {
				if j == posl[i] {
					continue
				} else if j < posl[i] {
					newImg.Set(i, j, img.At(i, j))
					newMask.Set(i, j, imgMask.At(i, j))
				} else {
					newImg.Set(i, j-1, img.At(i, j))
					newMask.Set(i, j-1, imgMask.At(i, j))
				}
			}
		}

		return newImg, retp, mnVal, newMask
	} else {
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

		return newImg, retp, mnVal, nil
	}

}

func carvingX(img image.Image, f func(image.Image, int) [][]float64, posp *[][]int, imgMask image.Image) (image.Image, *[][]int, float64, image.Image) {
	bounds := img.Bounds()

	// transpose
	dx := bounds.Dy()
	dy := bounds.Dx()

	energy := f(img, 1)
	setMask(&energy, imgMask, dy, dx)

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

	for i := dx - 1; i >= 0; i-- {
		posl[i] = pos
		pos = pre[i][pos]
	}

	newBounds := image.Rectangle{image.Pt(0, 0), image.Pt(dy-1, dx)}
	newImg := image.NewRGBA(newBounds)

	retp := (*[][]int)(nil)

	if posp != nil {
		posa := *posp
		reta := make([][]int, dy-1)
		for i := 0; i < dy-1; i++ {
			reta[i] = make([]int, dx)
		}
		retp = &reta

		for i := 0; i < dx; i++ {
			for j := 0; j < dy; j++ {
				if j == posl[i] {
					continue
				} else if j < posl[i] {
					reta[j][i] = posa[j][i]
				} else {
					reta[j-1][i] = posa[j][i]
				}
			}
		}
	}

	if imgMask != nil {
		newMask := image.NewRGBA(newBounds)

		for i := 0; i < dx; i++ {
			for j := 0; j < dy; j++ {
				if j == posl[i] {
					continue
				} else if j < posl[i] {
					newMask.Set(j, i, imgMask.At(j, i))
					newImg.Set(j, i, img.At(j, i))
				} else {
					newMask.Set(j-1, i, imgMask.At(j, i))
					newImg.Set(j-1, i, img.At(j, i))
				}
			}
		}

		return newImg, retp, mnVal, newMask
	} else {

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
		return newImg, retp, mnVal, nil
	}

}
func enlargeY(img image.Image, f func(image.Image, int) [][]float64, num int) (image.Image, image.Image) {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	energy := f(img, 2)

	dir := []int{-1, 0, 1}

	costflow := NewCostFlow(dx * dy * 2)

	id := func(x, y, z int) int {
		return (x*dy+y)*2 + z + 1
	}

	for i := 0; i < dy; i++ {
		costflow.AddEdge(costflow.S, id(0, i, 0), 1, 0)
		costflow.AddEdge(id(dx-1, i, 1), costflow.T, 1, 0)
	}

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			costflow.AddEdge(id(i, j, 0), id(i, j, 1), 1, energy[i][j])
			if i+1 < dx {
				for k := 0; k < 3; k++ {
					if j+dir[k] >= 0 && j+dir[k] < dy {
						costflow.AddEdge(id(i, j, 1), id(i+1, j+dir[k], 0), 1, 0)
					}
				}
			}
		}
	}
	costflow.MincostMaxflow(num)
	println("costflow done")

	newBounds := image.Rectangle{image.Pt(0, 0), image.Pt(dx, dy+num)}
	newImg := image.NewRGBA(newBounds)
	markImg := image.NewRGBA(newBounds)

	for i := 0; i < dx; i++ {
		offset := 0
		for j := 0; j < dy; j++ {
			newImg.Set(i, j+offset, img.At(i, j))
			markImg.Set(i, j+offset, img.At(i, j))
			if costflow.Check(id(i, j, 0), id(i, j, 1)) {
				offset++
				newImg.Set(i, j+offset, img.At(i, j))
				markImg.Set(i, j+offset, colornames.Red)
			}
		}
	}
	return newImg, markImg
}

func Resize(img image.Image, energyFunc func(image.Image, int) [][]float64, ndx, ndy int, imgMask image.Image) (image.Image, image.Image) {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	if ndx <= dx && ndy <= dy {
		if ndx == dx && ndy == dy {
			return img, img
		} else if ndx == dx {
			diff := dy - ndy
			bar := Pbar{}
			bar.Init(diff)
			position := make([][]int, dx)
			for i := 0; i < dx; i++ {
				position[i] = make([]int, dy)
				for j := 0; j < dy; j++ {
					position[i][j] = i*dy + j
				}
			}
			posp := &position
			oriImg := img
			if imgMask == nil {
				for i := 0; i < diff; i++ {
					img, posp, _, _ = carvingY(img, energyFunc, posp, imgMask)
					bar.Tick()
				}
			} else {
				for i := 0; i < diff; i++ {
					img, posp, _, imgMask = carvingY(img, energyFunc, posp, imgMask)
					bar.Tick()
				}
			}
			return img, mark(oriImg, posp)
		} else if ndy == dy {
			fImg := flipImg(img)
			fMask := (image.Image)(nil)
			if imgMask != nil {
				fMask = flipImg(imgMask)
			}

			retImg, retImg2 := Resize(fImg, energyFunc, ndy, ndx, fMask)
			return flipImg(retImg), flipImg(retImg2)
		} else {
			diffx := dx - ndx + 1
			diffy := dy - ndy + 1

			if diffx < diffy {
				fImg := flipImg(img)
				fMask := (image.Image)(nil)
				if imgMask != nil {
					fMask = flipImg(imgMask)
				}
				retImg, retImg2 := Resize(fImg, energyFunc, ndy, ndx, fMask)
				return flipImg(retImg), flipImg(retImg2)
			} else {
				println("cut two sides")
				I := make([][]image.Image, 2)
				I[0] = make([]image.Image, diffy)
				I[1] = make([]image.Image, diffy)
				Mask := make([][]image.Image, 2)
				Mask[0] = make([]image.Image, diffy)
				Mask[1] = make([]image.Image, diffy)
				T := make([][]float64, 2)
				T[0] = make([]float64, diffy)
				T[1] = make([]float64, diffy)
				from := make([][]int, diffx)
				for i := 0; i < diffx; i++ {
					from[i] = make([]int, diffy)
				}
				I[0][0] = img
				T[0][0] = 0
				bar := Pbar{}
				bar.Init(diffx * diffy)
				println("start", diffx, diffy)
				for i := 0; i < diffx; i++ {
					for j := 0; j < diffy; j++ {
						println(i, j)
						if i == 0 && j == 0 {
							continue
						}
						T[i&1][j] = math.Inf(1)
						if i > 0 {
							img, _, val, nmask := carvingX(I[i&1^1][j], energyFunc, nil, Mask[i&1^1][j])
							if val+T[i&1^1][j] < T[i&1][j] {
								T[i&1][j] = val + T[i&1^1][j]
								I[i&1][j] = img
								Mask[i&1][j] = nmask
								from[i][j] = 0
							}
						}
						if j > 0 {
							img, _, val, nmask := carvingY(I[i&1][j-1], energyFunc, nil, Mask[i&1][j-1])
							if val+T[i&1][j-1] < T[i&1][j] {
								T[i&1][j] = val + T[i&1][j-1]
								I[i&1][j] = img
								Mask[i&1][j] = nmask
								from[i][j] = 1
							}
						}
						bar.Tick()
					}
				}
				steps := make([]int, diffx+diffy)
				cnt := 0
				tx := diffx - 1
				ty := diffy - 1
				for tx != 0 || ty != 0 {
					steps[cnt] = from[tx][ty]
					cnt++
					if from[tx][ty] == 0 {
						tx--
					} else {
						ty--
					}
				}
				position := make([][]int, dx)
				for i := 0; i < dx; i++ {
					position[i] = make([]int, dy)
					for j := 0; j < dy; j++ {
						position[i][j] = i*dy + j
					}
				}
				posp := &position
				nowImg := img
				if imgMask == nil {
					for i := cnt - 1; i >= 0; i-- {
						println(i, "/", cnt, " : ", steps[i])
						if steps[i] == 0 {
							nowImg, posp, _, _ = carvingX(nowImg, energyFunc, posp, imgMask)
						} else {
							nowImg, posp, _, _ = carvingY(nowImg, energyFunc, posp, imgMask)
						}
					}
				} else {
					for i := cnt - 1; i >= 0; i-- {
						println(i, "/", cnt, " : ", steps[i])
						if steps[i] == 0 {
							nowImg, posp, _, imgMask = carvingX(nowImg, energyFunc, posp, imgMask)
						} else {
							nowImg, posp, _, imgMask = carvingY(nowImg, energyFunc, posp, imgMask)
						}
					}
				}

				return nowImg, mark(img, posp)
			}
		}
	} else {
		if ndx == dx {
			// enlarge y
			diff := ndy - dy
			img, img2 := enlargeY(img, energyFunc, diff)
			return img, img2
		} else if ndy == dy {
			// enlarge x

			// todo
			fImg := flipImg(img)
			fMask := (image.Image)(nil)
			if imgMask != nil {
				fMask = flipImg(imgMask)
			}

			retImg, retImg2 := Resize(fImg, energyFunc, ndy, ndx, fMask)
			return flipImg(retImg), flipImg(retImg2)
		} else {
			panic(errors.New("cannot enlarge both sides yet"))
		}
		return img, img
	}
}
