# SC

[![GoDoc](https://godoc.org/github.com/zrt/SC?status.svg)](https://godoc.org/github.com/zrt/SC)
[![Go Report Card](https://goreportcard.com/badge/github.com/zrt/SC)](https://goreportcard.com/report/github.com/zrt/SC)
[![License](https://img.shields.io/badge/LICENSE-GLWTPL-green.svg)](https://github.com/zrt/SC/blob/master/LICENSE)

A [Seam Carving algorithm](https://en.wikipedia.org/wiki/Seam_carving) implementation in Go with:
- [x]  CPU-only
- [ ] Uses all CPU cores in parallel
- [x] Supports PNG, JPEG files
- [x] Supports reduce image size
- [x] Supports increase image size
- [x] Supports region protection (imgMask green)
- [x] Supports region erasure (imgMask red)
- [ ] Polish API


## Usage

```bash
go get -u github.com/zrt/SC
```


```go
import "github.com/zrt/SC"
````

## Example

```go
package main

import . "github.com/zrt/SC"

func main() {

	img := LoadJPEG("input.jpg")

	energyFunc := L2NormEnergyFunc
	energyImg := SliceToImg(energyFunc(img, 1+2))
	// show energy img
	SavePNG(energyImg, "img_energy.png")

	newImg, newImgSeam := Resize(img, energyFunc, 1280*8/10, 868*8/10, nil)
	SavePNG(newImg, "output.png")
	SavePNG(newImgSeam, "output_seam.png")
}
```

### input.jpg

![example](https://github.com/zrt/SC/blob/master/_example/example1/input.jpg)

### img_energy.png

![example](https://github.com/zrt/SC/blob/master/_example/example1/img_energy.png)

### output.png

![example](https://github.com/zrt/SC/blob/master/_example/example1/output.png)

### output_seam.png

![example](https://github.com/zrt/SC/blob/master/_example/example1/output_seam.png)

### increase.png

![example](https://github.com/zrt/SC/blob/master/_example/example1/increase.png)

### increase_seam.png

![example](https://github.com/zrt/SC/blob/master/_example/example1/increase_seam.png)


## Links

- [seam carving wiki](https://en.wikipedia.org/wiki/Seam_carving)

## License

This project is licensed under [GLWTPL](https://github.com/me-shaon/GLWTPL).

