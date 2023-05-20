package main

import "fmt"

func heapSort(array []int) []int {
	for i := len(array) / 2; i >= 0; i-- {
		heapify(array, i, len(array))
	}
	for i := len(array) - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, 0, i)
	}
	return array

}

func heapify(array []int, IndicePai int, max int) {
	for {
		indiceFilhoEsq := 2*IndicePai + 1
		indiceFilhoDir := 2*IndicePai + 2
		j := IndicePai
		if (indiceFilhoEsq < max) && (array[IndicePai] < array[indiceFilhoEsq]) {
			IndicePai = indiceFilhoEsq
		}
		if (indiceFilhoDir < max) && (array[IndicePai] < array[indiceFilhoDir]) {
			IndicePai = indiceFilhoDir
		}
		if IndicePai == j {
			break
		}
		array[j], array[IndicePai] = array[IndicePai], array[j]
	}
}

func main() {
	array := []int{5, 3, 2, 1, 4} //, 30, 12, 43, 34454, 4353, 453, 4221, 56, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Print(heapSort(array))
}
