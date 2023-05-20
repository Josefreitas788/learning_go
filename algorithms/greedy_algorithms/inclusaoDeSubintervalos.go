package main

type intervalo struct {
	x1 int
	x2 int
}

func SortAlimento(a []intervalo) []intervalo {
	if len(a) <= 1 {
		return a
	}

	meio := len(a) / 2
	esquerda := SortAlimento(a[:meio])
	direita := SortAlimento(a[meio:])

	return mergeAlimento(esquerda, direita)
}

func mergeAlimento(esquerda, direita []intervalo) []intervalo {
	tamanho := len(esquerda) + len(direita)
	resultado := make([]intervalo, tamanho)
	for i := 0; i < tamanho; i++ {
		//TODO implementar
	}

	return resultado

}

func inclusao_subintervalos(intevalos []intervalo, intervalo intervalo) []intervalo {

}
