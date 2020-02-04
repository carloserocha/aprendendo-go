// Pacote principal
package main

// bibliotecas chamadas
import (
	"fmt"
	"time"
)

func main() {
	// Println - imprimir com quebra de linha
	// Printf - imprimir com tratamento para numbers e strings, especificando qual. Não esqueça do \n para quebra de linha.
	// Print - simplemente imprimir sem quebra de linha

	/*
		Printf:
			%T: Tipagem
			%s: Valor
			%d: númerico de base 10
	*/

	var firstName = "Ken"
	var lastName = "Thompson"
	var years = time.Now().Year() - 1943

	fmt.Println(firstName, lastName)
	fmt.Printf("%s %s\n", firstName, lastName)
	fmt.Printf("%s %s is %d years old\n", firstName, lastName, years)

	fmt.Println("Say goodbye")
}
