package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {

	arr := []int{5, 3, 2, 1, 4, 30, 12, 43, 34454, 4353, 453, 4221, 56, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Print(bubbleSort(arr))
}
