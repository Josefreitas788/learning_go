package main

import "fmt"

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i
		for j > 0 && array[j-1] > array[j] {
			array[j-1], array[j] = array[j], array[j-1]
			j--
		}
	}
	return array
}

func main() {

	array := []int{5, 3, 2, 1, 4, 30, 12, 43, 34454, 4353, 453, 4221, 56, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Print(insertionSort(array))
}
