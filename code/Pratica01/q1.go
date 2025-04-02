package main

import "fmt"

func main() {
	var num int

	fmt.Println("Escreva um número.")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("O número é positivo")

	} else if num < 0 {
		fmt.Println("O número é negativo")

	} else {
		fmt.Println("O número é nulo")
	}

}
