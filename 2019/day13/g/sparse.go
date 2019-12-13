package g

type SparseImage struct {
	Back Color
	Data map[Vector]Color
}

func NewSparseImage(back Color) *SparseImage {
	return &SparseImage{
		Back: back,
		Data: map[Vector]Color{},
	}
}

func (m *SparseImage) MinMax() (min, max Vector) {
	for px := range m.Data {
		min = min.Min(px)
		max = max.Max(px)
	}
	return min, max
}

func (m *SparseImage) At(at Vector) Color {
	color, ok := m.Data[at]
	if !ok {
		return m.Back
	}
	return color
}

func (m *SparseImage) Set(at Vector, color Color) {
	m.Data[at] = color
}

func (m *SparseImage) Image() *Image {
	min, max := m.MinMax()
	size := max.Sub(min).Add(Vector{1, 1})

	n := NewImage(size)
	for px, color := range m.Data {
		n.Set(px, color)
	}

	return n
}

func (m *SparseImage) Count(b Color) int64 {
	var count int64
	for _, v := range m.Data {
		if v == b {
			count++
		}
	}
	return count
}
