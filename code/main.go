package main

import "fmt"

func x(arr []int)

func main() {

	// subarrays
	arr := []int{1, -1, 2, -3, -4, 5, -5, 10, 1}
	l, r := 0, 0
	max := 0
	sum := 0
	ml, mr := 0, 0
	for r < len(arr) {
		sum += arr[r]
		fmt.Println("l", l, "r", r, "sum", sum, "max", max)
		if sum <= 0 {
			sum = 0
			l = r + 1
			r = l
		} else {
			r++
		}
		if sum > max {
			max = sum
			ml, mr = l, r
		}

	}
	fmt.Println(arr[ml:mr], "max", max)

}
