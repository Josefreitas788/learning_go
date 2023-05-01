package main

import (
	"fmt"
)

func recursivebBinarySearch(array []int, value int) int {
	if len(array) == 0 {
		return -1
	}

	mid := len(array) / 2

	if array[mid] == value {
		return array[mid]
	}

	if array[mid] > value {
		return recursivebBinarySearch(array[:mid], value)
	}

	return recursivebBinarySearch(array[mid+1:], value)

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 16}
	value := 1
	index := recursivebBinarySearch(array, value)

	fmt.Println(index)

}
