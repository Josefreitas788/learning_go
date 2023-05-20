package main

import "fmt"

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result[l+r] = left[l]
			l++
		} else {
			result[l+r] = right[r]
			r++
		}
	}
	for l < len(left) {
		result[l+r] = left[l]
		l++
	}
	for r < len(right) {
		result[l+r] = right[r]
		r++
	}
	return result
}

func main() {
	arr := []int{5, 3, 2, 1, 4, 30, 12, 43, 34454, 4353, 453, 4221, 56, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Print(mergeSort(arr))
}
