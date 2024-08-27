package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 1, 1, 2, 2, 2, 5}
	k := 3

	bits := [32]int{}
	for i := range arr {
		// increment bits[i] count if bit is set in arr[i]
	}

	res := float64(0)
	for i := 0; i < 32; i++ {
		if bits[i]%k == 1 {
			// that bit is 1
			res += math.Pow(2, float64(i))
		}
		// else bit is 0 // no need to add since its power will be 0
	}
	fmt.Println(res)
}
