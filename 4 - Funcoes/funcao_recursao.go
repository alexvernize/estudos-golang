package main

import "fmt"

/*
RECURSÃO EM GO

Recursão é quando uma função chama a si mesma.
É útil para problemas que podem ser divididos em subproblemas similares.
*/

// fatorial calcula o fatorial de n recursivamente
func fatorial(n int) int {
	// Caso base - interrompe a recursão
	if n == 0 {
		return 1
	}

	// Passo recursivo - chama a si mesma com n-1
	return n * fatorial(n-1)
}

/*
FIBONACCI RECURSIVO

A sequência de Fibonacci é um clássico exemplo de recursão.
Porém, esta implementação é ineficiente para valores grandes.
*/
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// FATORIAL RECURSIVO
	fmt.Println("Fatorial de 5:", fatorial(5)) // 120

	// FIBONACCI RECURSIVO
	fmt.Println("Fibonacci (n=10):", fibonacci(10)) // 55

	/*
	   CONSIDERAÇÕES SOBRE RECURSÃO:
	   1. Sempre deve ter um caso base para evitar loop infinito
	   2. Pode consumir muita memória com chamadas aninhadas
	   3. Em Go, a pilha de chamadas é limitada (cuidado com recursão profunda)
	   4. Alguns problemas são naturalmente recursivos (árvores, divisão e conquista)
	*/
}
