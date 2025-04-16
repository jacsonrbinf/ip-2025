package main

import "fmt"

func main() {
	fmt.Println("Busca binaria...")
	list := []int{2, 3, 5, 7, 10, 13, 15, 19, 23, 30}
	posicao := buscaBinaria(list, 13)
	fmt.Println(posicao)
}
func buscaBinaria(list []int, alvo int) int {
	esquerdo := 0
	direito := len(list) - 1
	for esquerdo <= direito {
		meio := (direito + esquerdo) / 2
		elem_meio := list[meio]
		if elem_meio == alvo {
			return meio
		}
		if elem_meio > alvo {
			direito = meio - 1
		}
		if elem_meio < alvo {
			esquerdo = meio + 1
		}
	}
	return -1
}
