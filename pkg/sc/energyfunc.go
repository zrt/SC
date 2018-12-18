package sc

import (
	"image"
	"math"
)

func L2NormEnergyFunc(img image.Image) [][]float64 {
	bounds := img.Bounds()

	ret := make([][]float64, bounds.Dx())
	for i := 0; i < bounds.Dx(); i++ {
		ret[i] = make([]float64, bounds.Dy())
	}

	for i := bounds.Min.X + 1; i < bounds.Max.X-1; i++ {
		for j := bounds.Min.Y + 1; j < bounds.Max.Y-1; j++ {
			tmp := uint32(0)
			r1, g1, b1, _ := img.At(i-1, j).RGBA()
			r2, g2, b2, _ := img.At(i+1, j).RGBA()
			tmp += (r2-r1)*(r2-r1) + (g2-g1)*(g2-g1) + (b2-b1)*(b2-b1)
			r1, g1, b1, _ = img.At(i, j-1).RGBA()
			r2, g2, b2, _ = img.At(i, j+1).RGBA()
			tmp += (r2-r1)*(r2-r1) + (g2-g1)*(g2-g1) + (b2-b1)*(b2-b1)
			ret[i][j] = math.Sqrt(float64(tmp))
		}
	}

	return ret
}
