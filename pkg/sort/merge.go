package main

import (
	"fmt"
)

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	left := make([]int, n1)
	right := make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		right[j] = arr[m+1+j]
	}

	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = left[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = right[j]
		j++
		k++
	}
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := l + (r-l)/2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
