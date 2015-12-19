// +build ignore

package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	code := "iwrupvqb"
	for i := 1; i < 1<<30; i++ {
		s := fmt.Sprintf("%s%d", code, i)
		sum := md5.Sum([]byte(s))
		if sum[0] == 0 && sum[1] == 0 && sum[2] == 0 {
			fmt.Println(i)
		}
	}
}
