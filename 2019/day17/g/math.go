package g

func Sign(v int64) int64 {
	if v < 0 {
		return -1
	} else if v > 0 {
		return 1
	}

	return 0
}

func Abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
