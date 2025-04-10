/*
Função para retornar o maior numero
*/
package main

import "fmt"

func main() {

	fmt.Println(numeromaior(20, 8, 10))

}
func numeromaior(x int, y int, z int) int {
	if x >= y && x >= z {
		return x
	} else if y >= x && y >= z {
		return y
	} else {
		return z
	}
}
