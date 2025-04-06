package main

import "fmt"

func main() {
	/*
	   FUNÇÕES ANÔNIMAS EM GO

	   São funções sem nome que podem ser declaradas e executadas imediatamente
	   ou atribuídas a variáveis. São úteis para:
	   - Implementações pequenas e locais
	   - Closures
	   - Callbacks
	*/

	// 1. Função anônima atribuída a variável
	quadrado := func(x int) int {
		return x * x
	}
	fmt.Println("Quadrado de 5:", quadrado(5)) // 25

	// 2. Função anônima auto-executável
	func() {
		fmt.Println("Executando função anônima imediatamente!")
	}() // Os parênteses no final executam a função

	// 3. Função anônima com parâmetros
	resultado := func(a, b int) int {
		return a + b
	}(10, 20) // Passando argumentos na execução
	fmt.Println("Resultado:", resultado) // 30

	// 4. Usando como closure
	contador := 0
	incrementar := func() int {
		contador++ // Captura a variável contador
		return contador
	}
	fmt.Println(incrementar()) // 1
	fmt.Println(incrementar()) // 2

	/*
	   DIFERENÇA ENTRE FUNÇÃO ANÔNIMA E CLOSURE:
	   - Toda closure é uma função anônima
	   - Mas nem toda função anônima é closure
	   - Closure captura variáveis do escopo externo
	*/
}
