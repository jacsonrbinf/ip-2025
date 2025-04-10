package main

import "fmt"

// Função para realizar uma busca sequencial
func buscaSequencial(array []int, valor int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == valor {
			return i // Retorna o índice do valor encontrado
		}
	}
	return -1 // Retorna -1 se o valor não for encontrado
}
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	valor := 5
	indice := buscaSequencial(array, valor)

	if indice != -1 {
		fmt.Printf("Valor encontrado no índice: %d\n", indice)
	} else {
		fmt.Println("Valor não encontrado no array.")
	}
}
