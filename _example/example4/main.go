package main

// example2
// by zrt

import (
	"fmt"
	. "github.com/zrt/SC"
	"image"
	"strings"

	"flag"
)

func main() {
	inputfile := flag.String("i", "", "input file name")
	outputfile := flag.String("o", "", "output file name")
	outputfile2 := flag.String("o2", "", "output mark file name")
	flag.Parse()
	if *inputfile == "" || *outputfile == "" || *outputfile2 == "" {
		panic("-h for usage")
	}
	img := image.Image(nil)
	if strings.HasSuffix(*inputfile, "png") {
		img = LoadPNG(*inputfile)
	} else {
		img = LoadJPEG(*inputfile)
	}
	fmt.Printf("%s %v\n", *inputfile, img.Bounds().Size()) // show img size

	energyFunc := L2NormEnergyFunc

	//newImg, newImg2 := Resize(img, energyFunc, 682-136, 680-136)
	xnewImg, xnewImg2 := Resize(img, energyFunc, img.Bounds().Dx()*20/10, img.Bounds().Dy(), nil)
	ynewImg, ynewImg2 := Resize(img, energyFunc, img.Bounds().Dx(), img.Bounds().Dy()*20/10, nil)

	fmt.Printf("%s %v\n", *outputfile, xnewImg.Bounds().Size()) // show newImg size
	SavePNG(xnewImg, "x_"+*outputfile)                          // show newImg
	SavePNG(xnewImg2, "x_"+*outputfile2)                        // show newImg
	fmt.Printf("%s %v\n", *outputfile, ynewImg.Bounds().Size()) // show newImg size
	SavePNG(ynewImg, "y_"+*outputfile)                          // show newImg
	SavePNG(ynewImg2, "y_"+*outputfile2)                        // show newImg

}
