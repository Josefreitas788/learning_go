package main

import "fmt"

func CountingSort(array []int) []int {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}

	count := make([]int, max+1)
	for _, v := range array {
		count[v]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	result := make([]int, len(array))
	for _, v := range array {
		result[count[v]-1] = v
		count[v]--
	}

	return result

}

func main() {
	// array := []int{5, 3, 2, 1, 4, 3, 7, 8, 2}
	array := []int{5, 4, 2, 1, 2}

	fmt.Print(CountingSort(array))
}
