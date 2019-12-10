package main

import "fmt"

func main() {
	layers := ParseLayers(Input, 25, 6)

	leastDense := layers[0]
	leastCount := layers[0].Count(0)
	for _, layer := range layers[1:] {
		count := layer.Count(0)
		if count < leastCount {
			leastDense = layer
			leastCount = count
		}
	}

	fmt.Println(leastDense.Count(1) * leastDense.Count(2))

	composite := NewImage(leastDense.Width, leastDense.Height)
	composite.Fill(Transparent)

	for i := len(layers) - 1; i >= 0; i-- {
		layers[i].DrawTo(composite)
	}

	composite.Print()
}

type Image struct {
	Data   []Color
	Width  int64
	Height int64
}

type Color = byte

const (
	Black       = Color(0)
	White       = Color(1)
	Transparent = Color(2)
)

func NewImage(width, height int64) *Image {
	return &Image{
		Data:   make([]byte, width*height),
		Width:  width,
		Height: height,
	}
}

func (image *Image) Fill(c Color) {
	for i := range image.Data {
		image.Data[i] = c
	}
}

func (image *Image) DrawTo(target *Image) {
	for i, src := range image.Data {
		if src != Transparent {
			target.Data[i] = src
		}
	}
}

func (image *Image) Print() {
	for y := int64(0); y < image.Height; y++ {
		for x := int64(0); x < image.Width; x++ {
			var c = image.Data[image.Index(x, y)]
			switch c {
			case Black:
				fmt.Print("█")
			case White:
				fmt.Print(" ")
			case Transparent:
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}

func (image *Image) Index(x, y int64) int64 {
	return image.Width*y + x
}

func (image *Image) Count(b Color) int64 {
	var count int64
	for _, v := range image.Data {
		if v == b {
			count++
		}
	}
	return count
}

func ParseLayers(input string, width, height int64) []*Image {
	var images []*Image
	for len(input) > 0 {
		image := NewImage(width, height)
		images = append(images, image)

		n := copy(image.Data, []byte(input))
		input = input[n:]

		for i, r := range image.Data {
			image.Data[i] = r - '0'
		}
	}

	return images
}
