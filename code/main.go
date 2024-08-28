package main

import "fmt"

func numberOfWay(n, m int) int {
	if n == 1 || m == 1 {
		return 1
	}
	return numberOfWay(n-1, m) + numberOfWay(n, m-1)
}

func main() {
	fmt.Println(numberOfWay(10, 10))
}
