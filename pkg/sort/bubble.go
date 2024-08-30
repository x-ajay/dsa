package main

import "fmt"

func bubble(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 5, 2, 3, 7, 2, 9}
	fmt.Println(bubble(arr))
}
