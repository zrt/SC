package main

// example1
// by zrt

import (
	"fmt"
	. "github.com/zrt/SeamCarvingGO/pkg/sc"
)

func main() {

	img := LoadJPEG("input.jpg")
	fmt.Printf("%v\n", img.Bounds().Size()) // show img size

	energyImg := SliceToImg(L2NormEnergyFunc(img))
	SavePNG(energyImg, "img_energy.png") // show energy img

	CarvingY(img, L2NormEnergyFunc, true) // show a seam example

	newImg := Resize(img, L2NormEnergyFunc, 230, 230)

	fmt.Printf("%v\n", newImg.Bounds().Size()) // show newImg size
	SavePNG(newImg, "output.png")              // show newImg

}
