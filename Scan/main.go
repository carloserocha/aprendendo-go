package main

import "fmt"

func main() {
	var name string

	fmt.Print("Qual seu nome? ")             // printa na tela a mensagem
	fmt.Scanf("%s", &name)                   // espera a entrada do usuário
	fmt.Printf("Olá %s, tudo bem? :)", name) // printa o nome
}
