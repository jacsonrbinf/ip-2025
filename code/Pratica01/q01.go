package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Print("Forneca o nome e a idade: ")
	fmt.Scan(&name, &age)
	fmt.Printf("Ola %s, voce tem %d ano(s) de idade.\n", name, age)
}
