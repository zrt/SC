package SC

import (
	"image"
	"image/color"
	"math"
)

func getFloat64Color(c color.Color) (float64, float64, float64, float64) {
	r, g, b, a := c.RGBA()
	return float64(r), float64(g), float64(b), float64(a)
}

func L2NormEnergyFunc(img image.Image, xy int) [][]float64 {
	bounds := img.Bounds()

	ret := make([][]float64, bounds.Dx())
	for i := 0; i < bounds.Dx(); i++ {
		ret[i] = make([]float64, bounds.Dy())
	}

	for i := bounds.Min.X + 1; i < bounds.Max.X-1; i++ {
		for j := bounds.Min.Y + 1; j < bounds.Max.Y-1; j++ {
			if xy&1 != 0 {
				r1, g1, b1, _ := getFloat64Color(img.At(i-1, j))
				r2, g2, b2, _ := getFloat64Color(img.At(i+1, j))
				tmp := (r2-r1)*(r2-r1) + (g2-g1)*(g2-g1) + (b2-b1)*(b2-b1)
				ret[i][j] += math.Sqrt(tmp)
			}
			if xy&2 != 0 {
				r1, g1, b1, _ := getFloat64Color(img.At(i, j-1))
				r2, g2, b2, _ := getFloat64Color(img.At(i, j+1))
				tmp := (r2-r1)*(r2-r1) + (g2-g1)*(g2-g1) + (b2-b1)*(b2-b1)
				ret[i][j] += math.Sqrt(tmp)
			}
		}
	}
	return ret
}

func L1NormEnergyFunc(img image.Image, xy int) [][]float64 {
	bounds := img.Bounds()

	ret := make([][]float64, bounds.Dx())
	for i := 0; i < bounds.Dx(); i++ {
		ret[i] = make([]float64, bounds.Dy())
	}

	for i := bounds.Min.X + 1; i < bounds.Max.X-1; i++ {
		for j := bounds.Min.Y + 1; j < bounds.Max.Y-1; j++ {
			if xy&1 != 0 {
				r1, g1, b1, _ := getFloat64Color(img.At(i-1, j))
				r2, g2, b2, _ := getFloat64Color(img.At(i+1, j))
				ret[i][j] += math.Abs(r2-r1) + math.Abs(g2-g1) + math.Abs(b2-b1)
			}
			if xy&2 != 0 {
				r1, g1, b1, _ := getFloat64Color(img.At(i, j-1))
				r2, g2, b2, _ := getFloat64Color(img.At(i, j+1))
				ret[i][j] += math.Abs(r2-r1) + math.Abs(g2-g1) + math.Abs(b2-b1)
			}

		}
	}
	return ret
}
