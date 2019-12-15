package g

import "fmt"

type Image struct {
	Data []Color
	Size Vector
}

type Color = byte

const (
	Black       = Color(0)
	White       = Color(1)
	Transparent = Color(2)
)

var DefaultColors = map[Color]rune{
	Black:       ' ',
	White:       '█',
	Transparent: '/',
}

func NewImage(size Vector) *Image {
	return &Image{
		Data: make([]byte, size.X*size.Y),
		Size: size,
	}
}

func (image *Image) Fill(c Color) {
	for i := range image.Data {
		image.Data[i] = c
	}
}

func (image *Image) Print(colors map[Color]rune) {
	for y := int64(0); y < image.Size.Y; y++ {
		for x := int64(0); x < image.Size.X; x++ {
			var c = image.Data[image.Index(Vector{x, y})]
			color, ok := colors[c]
			if ok {
				fmt.Print(string(color))
			} else {
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}

func (image *Image) Index(at Vector) int64 {
	return image.Size.X*at.Y + at.X
}

func (image *Image) Set(at Vector, c Color) {
	image.Data[image.Index(at)] = c
}

func (image *Image) At(at Vector) Color {
	return image.Data[image.Index(at)]
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
