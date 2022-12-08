// +build ignore

package main

import "fmt"

type Cookie struct {
	Name     string
	Taste    Taste
	Calories int
}

type Taste [4]int

func (a *Taste) AddScale(b *Taste, amount int) {
	for i := range a {
		a[i] += b[i] * amount
	}
}

func pos(v int) int {
	if v < 0 {
		return 0
	}
	return v
}

func (t *Taste) Score() (score int) {
	score = 1
	for _, v := range t {
		if v <= 0 {
			return 0
		}
		score *= v
	}
	return score
}

var withcalories bool

func Score(ratio []int, cookies []Cookie) int {
	total := Taste{}
	calories := 0
	for i := range cookies {
		total.AddScale(&cookies[i].Taste, ratio[i])
		calories += cookies[i].Calories * ratio[i]
	}
	if withcalories && calories != 500 {
		return 0
	}
	return total.Score()
}

func all(index int, left int, ratio []int, cookies []Cookie) int {
	if index == len(cookies)-1 {
		ratio[index] = left
		return Score(ratio, cookies)
	}

	best := 0
	for amount := 0; amount < left; amount++ {
		ratio[index] = amount
		score := all(index+1, left-amount, ratio, cookies)
		if best < score {
			best = score
		}
	}
	return best
}

func BestTaste(cookies []Cookie) int {
	ratios := make([]int, len(cookies))
	return all(0, 100, ratios, cookies)
}

func main() {
	test := []Cookie{
		{"Butterscotch", Taste{-1, -2, 6, 3}, 8},
		{"Cinnamon", Taste{2, 3, -2, -1}, 3},
	}
	fmt.Println(BestTaste(test))
	fmt.Println(Score([]int{44, 56}, test))

	cookies := []Cookie{
		{"Sprinkles", Taste{5, -1, 0, 0}, 5},
		{"PeanutButter", Taste{-1, 3, 0, 0}, 1},
		{"Frosting", Taste{0, -1, 4, 0}, 6},
		{"Sugar", Taste{-1, 0, 0, 2}, 8},
	}
	fmt.Println(cookies)
	fmt.Println(BestTaste(cookies))
	withcalories = true
	fmt.Println(Score([]int{44, 56}, test))
	fmt.Println(Score([]int{40, 60}, test))
	fmt.Println(BestTaste(cookies))
}
