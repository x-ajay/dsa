Here's how you can implement the sorting and searching algorithms in Go (Golang).

## Sorting Algorithms

### 1. Bubble Sort

**Description:**

Bubble Sort is a simple algorithm that works by repeatedly swapping adjacent elements if they are in the wrong order. The process is repeated until the entire list is sorted.

**Go Implementation:**

```go
package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
```

### 2. Insertion Sort

**Description:**

Insertion Sort builds the sorted array one element at a time. It repeatedly picks the next element and inserts it into the correct position in the already sorted part of the array.

**Go Implementation:**

```go
package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
```

### 3. Selection Sort

**Description:**

Selection Sort works by repeatedly selecting the smallest element from the unsorted portion of the array and swapping it with the first unsorted element.

**Go Implementation:**

```go
package main

import (
	"fmt"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	selectionSort(arr)
	fmt.Println("Sorted array:", arr)
}
```

### 4. Quick Sort

**Description:**

Quick Sort is a divide-and-conquer algorithm that works by selecting a "pivot" element from the array and partitioning the other elements into two sub-arrays according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively.

**Go Implementation:**

```go
package main

import (
	"fmt"
)

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Sorted array:", arr)
}
```

### 5. Merge Sort

**Description:**

Merge Sort is a divide-and-conquer algorithm that divides the array into two halves, sorts each half, and then merges the two sorted halves back together.

**Go Implementation:**

```go
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
```

## Searching Algorithms

### 1. Linear Search

**Description:**

Linear Search is the simplest search algorithm. It sequentially checks each element of the array until it finds the target value.

**Go Implementation:**

```go
package main

import (
	"fmt"
)

func linearSearch(arr []int, x int) int {
	for i, v := range arr {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 10
	result := linearSearch(arr, x)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
}
```

### 2. Binary Search

**Description:**

Binary Search is an efficient algorithm for finding an element in a sorted array. It works by repeatedly dividing the search interval in half.

**Go Implementation:**

```go
package main

import (
	"fmt"
)

func binarySearch(arr []int, x int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		m := l + (r-l)/2

		if arr[m] == x {
			return m
		}

		if arr[m] < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 10
	result := binarySearch(arr, x)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
}
```

## Conclusion

These sorting and searching algorithms form the foundation of many more complex algorithms and data structures. Understanding how they work and how to implement them in Go can greatly enhance your problem-solving skills in computer science and software development.

---

In competitive programming, sorting is a common task that often requires efficiency and performance. Go provides efficient built-in sorting functions in the `sort` package, which are suitable for most competitive programming tasks.

## Sorting Techniques in Go for Competitive Programming

### 1. Built-in Sort Functions

The `sort` package in Go provides built-in functions to sort slices of integers, floats, and strings. These functions are optimized and should be your go-to choice in competitive programming.

#### Example: Sorting a Slice of Integers

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	sort.Ints(arr)
	fmt.Println("Sorted array:", arr)
}
```

- `sort.Ints(arr)`: Sorts a slice of integers in ascending order.
- `sort.Float64s(arr)`: Sorts a slice of floats in ascending order.
- `sort.Strings(arr)`: Sorts a slice of strings in ascending order.

### 2. Sorting in Descending Order

The built-in sorting functions sort in ascending order by default. To sort in descending order, you can use `sort.Sort()` with a custom comparator or simply reverse the slice after sorting.

#### Example: Sorting Integers in Descending Order

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println("Sorted array in descending order:", arr)
}
```

### 3. Custom Sorting

For more complex sorting needs, where you need to sort based on multiple criteria or custom rules, you can define your own sorting logic using `sort.Slice()`.

#### Example: Custom Sorting with Multiple Criteria

Suppose you have a slice of structs, and you want to sort based on multiple fields:

```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"Dave", 20},
	}

	// Sort by Age, then by Name
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Name < people[j].Name
		}
		return people[i].Age < people[j].Age
	})

	fmt.Println("Sorted people:", people)
}
```
