package main

import (
	"fmt"
	"math"
)

func isPrime(n int) {
	// seive of erathosthesis
	arr := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		arr[i] = true
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		for j := i + i; j <= n; j += i {
			arr[j] = false
		}

	}
	for i := 0; i < len(arr); i++ {
		if arr[i] {
			fmt.Println(i)
		}
	}
}

func main() {
	isPrime(10)
}
