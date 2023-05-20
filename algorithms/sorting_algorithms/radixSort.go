package main

import "fmt"

func radixSort(array []int) []int {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		countingSort(array, exp)
	}

	return array
}

func countingSort(array []int, exp int) {
	n := len(array)
	count := make([]int, 10)
	result := make([]int, n)

	for _, v := range array {
		count[(v/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		result[count[(array[i]/exp)%10]-1] = array[i]
		count[(array[i]/exp)%10]--
	}

	for i := 0; i < n; i++ {
		array[i] = result[i]
	}
}

func main() {
	array := []int{5, 3, 2, 1, 4, 30, 12, 43, 34454, 4353, 453, 4221, 56, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Print(radixSort(array))
}
