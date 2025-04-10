/*
 função que divida dois números e retorne dois valores:
 Um valor float64 com o resultado da divisão e outro valor booleano
 indicando se houve erro na divisão, caso o divisor seja 0 (zero):
*/

package main

import f "fmt"

func main() {
	// Declaração e leitura dos números
	n1, n2 := 0.0, 0.0
	f.Print("Digite o primeiro número: ")
	f.Scan(&n1)
	f.Print("Digite o segundo número: ")
	f.Scan(&n2)
	// Cálculo da divisao através da chamada à função
	e, r := divisao(n1, n2) // Cálculo da divisao através da função
	if !e {
		f.Printf("Resultado = %.2f\n", r)
	} else {
		f.Printf("Divisao nao permitida\n")
	}

}

// Definição da função div
func divisao(numero1, numero2 float64) (bool, float64) {
	erro := false
	result := 0.0
	if numero2 == 0 {
		erro = true
	} else {
		result = numero1 / numero1
	}
	return erro, result
}
