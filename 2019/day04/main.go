package main

import "fmt"

func main() {
	var countAdjacent int64
	var countCluster int64
	for p := 138307; p <= 654504; p++ {
		pass := NewPassword(int64(p))

		if pass.AdjacentTwo() && pass.Increasing() {
			countAdjacent++
		}

		if pass.ClusterTwo() && pass.Increasing() {
			countCluster++
		}
	}

	fmt.Println("count adjacent", countAdjacent)
	fmt.Println("count cluster", countCluster)
}

type Password [6]byte

func NewPassword(p int64) (pass Password) {
	pass[0] = byte((p / 100000) % 10)
	pass[1] = byte((p / 10000) % 10)
	pass[2] = byte((p / 1000) % 10)
	pass[3] = byte((p / 100) % 10)
	pass[4] = byte((p / 10) % 10)
	pass[5] = byte((p / 1) % 10)
	return pass
}

func (pass Password) Int64() int64 {
	return int64(pass[0])*100000 +
		int64(pass[1])*10000 +
		int64(pass[2])*1000 +
		int64(pass[3])*100 +
		int64(pass[4])*10 +
		int64(pass[5])*1
}

func (pass Password) InRange(min, max int64) bool {
	v := pass.Int64()
	return min <= v && v <= max
}

func (pass Password) AdjacentTwo() bool {
	return pass[0] == pass[1] ||
		pass[1] == pass[2] ||
		pass[2] == pass[3] ||
		pass[3] == pass[4] ||
		pass[4] == pass[5]
}

func (pass Password) ClusterTwo() bool {
	for i := 0; i < len(pass); {
		x := pass[i]
		leading := countLeading(x, pass[i:])

		if leading == 2 {
			return true
		}
		if leading == 0 {
			break
		}

		i += int(leading)
	}
	return false
}

func countLeading(z byte, xs []byte) int64 {
	for i, x := range xs {
		if x != z {
			return int64(i)
		}
	}
	return int64(len(xs))
}

func (pass Password) Increasing() bool {
	return pass[0] <= pass[1] &&
		pass[1] <= pass[2] &&
		pass[2] <= pass[3] &&
		pass[3] <= pass[4] &&
		pass[4] <= pass[5]
}
