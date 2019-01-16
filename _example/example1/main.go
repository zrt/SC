package main

// example1
// by zrt

import (
	"fmt"
	. "github.com/zrt/SC"
)

func main() {

	img := LoadJPEG("input.jpg")
	fmt.Printf("%v\n", img.Bounds().Size()) // show img size

	energyFunc := L2NormEnergyFunc
	energyImg := SliceToImg(energyFunc(img, 1+2))
	SavePNG(energyImg, "img_energy.png") // show energy img

	for i := 0; i < 10; i++ {
		newImg, newImgWithSeam := Resize(img, energyFunc, 1280*(i+1)/10, 868, nil)
		fmt.Printf("%v\n", newImg.Bounds().Size())                   // show newImg size
		SavePNG(newImg, fmt.Sprintf("output%d.png", i))              // show newImg
		SavePNG(newImgWithSeam, fmt.Sprintf("output%d_seam.png", i)) // show newImg
	}

}
