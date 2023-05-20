package main

import (
	"fmt"
)

func caminhoParaCasaDaNamorada(distancias []int, autonomia int) int {
	tanqueCheio := autonomia
	paradas := 0
	for i := 0; i < len(distancias)-1; i++ {
		if distancias[i]+distancias[i+1] <= autonomia {
			autonomia -= distancias[i]
		} else {
			autonomia = tanqueCheio
			paradas++
		}
	}
	return paradas
}

func main() {
	distancias := []int{10, 20, 30, 40}
	autonomia := 30
	fmt.Println(caminhoParaCasaDaNamorada(distancias, autonomia))
}
