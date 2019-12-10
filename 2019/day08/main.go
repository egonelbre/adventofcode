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
}

type Image struct {
	Data   []byte
	Width  int64
	Height int64
}

func NewImage(width, height int64) *Image {
	return &Image{
		Data: make([]byte, width*height),
	}
}

func (image *Image) Index(x, y int64) int64 {
	return image.Height*y + x
}

func (image *Image) Count(b byte) int64 {
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
