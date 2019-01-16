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

	energyFunc := L2NormEnergyFunc
	energyImg := SliceToImg(energyFunc(img, 1+2))
	SavePNG(energyImg, "img_energy.png") // show energy img

	//newImg, newImg2 := Resize(img, energyFunc, 682-136, 680-136)
	newImg, newImg2 := Resize(img, energyFunc, img.Bounds().Dx()*8/10, img.Bounds().Dy()*8/10, nil)

	fmt.Printf("%v\n", newImg.Bounds().Size()) // show newImg size
	SavePNG(newImg, "output0.2.png")           // show newImg
	SavePNG(newImg2, "output0.2.mark.png")     // show newImg
}
