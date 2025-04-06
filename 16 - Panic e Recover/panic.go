package main

import "fmt"

func main() {
	/*
	   PANIC EM GO

	   - `panic` interrompe o fluxo normal do programa e inicia o "pânico".
	   - Quando um panic ocorre, a função atual para de executar e começa a desempilhar as chamadas.
	   - Se nenhum `recover` for encontrado, o programa termina com uma mensagem de erro.
	*/

	fmt.Println("Início do programa")

	// Causando um panic
	panic("Algo muito errado aconteceu!") // Mensagem de panic

	// Este código NUNCA será executado
	fmt.Println("Fim do programa (não será exibido)")
}
