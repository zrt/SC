package main

// example2
// by zrt

import (
	"fmt"
	. "github.com/zrt/SC"
)

func main() {

	img := LoadJPEG("input.jpg")
	fmt.Printf("%v\n", img.Bounds().Size()) // show img size
	imgMask := LoadPNG("input_mask.png")
	fmt.Printf("%v\n", imgMask.Bounds().Size()) // show img size

	energyFunc := L2NormEnergyFunc
	newImg, newImg2 := Resize(img, energyFunc, img.Bounds().Dx()*95/100, img.Bounds().Dy()*95/100, imgMask)

	fmt.Printf("%v\n", newImg.Bounds().Size()) // show newImg size
	SavePNG(newImg, "output0.2.png")           // show newImg
	SavePNG(newImg2, "output0.2.mark.png")     // show newImg

}
