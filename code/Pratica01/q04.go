/*
Programa para calcular a media global das notas dos alunos de uma disciplina.
*/

package main

import (
	"fmt"
)

func main() {
	var numAlunos int
	fmt.Print("Digite o número de alunos: ")
	fmt.Scan(&numAlunos)

	if numAlunos <= 0 {
		fmt.Println("Número inválido de alunos.")
		return
	}

	var soma, notas float64

	for i := 0; i < numAlunos; i++ {
		fmt.Printf("Digite a nota do aluno %d: ", i+1)
		fmt.Scan(&notas)
		soma += notas
	}

	mediaGlobal := soma / float64(numAlunos)
	fmt.Printf("A média global da turma é: %.2f\n", mediaGlobal)
}
