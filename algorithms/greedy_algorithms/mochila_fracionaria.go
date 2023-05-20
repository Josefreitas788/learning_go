package main

import "fmt"

type Alimento struct {
	gramaPorAlimento      float32
	satisfacaoPorAlimento float32
}

// I used the merge sort algorithm to sort the food because the algorithm used by go for sorting slices is the quickSort(n^2) and i wanted a (n log n) algorithm
func SortAlimento(a []Alimento) []Alimento {
	if len(a) <= 1 {
		return a
	}

	meio := len(a) / 2
	esquerda := SortAlimento(a[:meio])
	direita := SortAlimento(a[meio:])

	return mergeAlimento(esquerda, direita)
}

func mergeAlimento(esquerda, direita []Alimento) []Alimento {
	tamanho := len(esquerda) + len(direita)
	i, j := 0, 0
	var resultado []Alimento

	for k := 0; k < tamanho; k++ {
		if esquerda[i].satisfacaoPorAlimento > direita[j].satisfacaoPorAlimento {
			resultado = append(resultado, esquerda[i])
			i++
		} else {
			resultado = append(resultado, direita[j])
			j++
		}
		if i == len(esquerda) {
			resultado = append(resultado, direita[j:]...)
			break
		}
		if j == len(direita) {
			resultado = append(resultado, esquerda[i:]...)
			break
		}
	}

	return resultado
}

func satisfaçãoMaxima(limiteTotal float32, gramaPorAlimento []float32, satisfacaoPorAlimento []float32) float32 {
	satisfacaoMaxima := float32(0)
	var alimento []Alimento

	for i := 0; i < len(gramaPorAlimento); i++ { //Theta(n)
		alimento = append(alimento, Alimento{gramaPorAlimento[i], satisfacaoPorAlimento[i]})
	}

	alimento = SortAlimento(alimento) //Theta(n log n)

	for i := 0; i < len(alimento); i++ { //Theta(n)
		if limiteTotal > 0 {
			if alimento[i].gramaPorAlimento <= limiteTotal {
				satisfacaoMaxima += alimento[i].satisfacaoPorAlimento * alimento[i].gramaPorAlimento
				limiteTotal -= alimento[i].gramaPorAlimento
			} else {
				satisfacaoMaxima += alimento[i].satisfacaoPorAlimento * limiteTotal / alimento[i].gramaPorAlimento
				limiteTotal = 0
				break
			}
		}
	}

	return satisfacaoMaxima
}

func main() {

	limiteTotal := float32(30.23)
	gramaPorAlimento := []float32{12, 2.4, 3.5, 4.5, 5.6, 6.7, 21.8, 8.9, 23.0, 10.1, 11.2, 12.3}
	satisfacaoPorAlimento := []float32{12, 2.4, 3.5, 4.5, 5.6, 6.7, 8.8, 8.9, 9.0, 0.1, 11.2, 16.3}

	resultado := satisfaçãoMaxima(limiteTotal, gramaPorAlimento, satisfacaoPorAlimento)
	fmt.Printf("%.4f\n", resultado)
}
