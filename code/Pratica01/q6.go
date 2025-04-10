/*
Função de nome soma que receba 2 valores e retorne a soma desses valores.
No programa principal leia dois números e calcule sua soma através da função criada.
*/

package main

import f "fmt"

func main() {
	// Declaração e leitura dos números a serem somados
	n1, n2 := 0, 0
	f.Print("Digite o primeiro número: ")
	f.Scan(&n1)
	f.Print("Digite o segundo número: ")
	f.Scan(&n2)
	// Cálculo da soma através da chamada à função
	s := soma(n1, n2) // Cálculo da soma através da função
	f.Printf("A soma entre %d e %d é %d\n", n1, n2, s)
}

// Definição da função soma
func soma(numero1, numero2 int) int {
	s := numero1 + numero2
	return s
}
