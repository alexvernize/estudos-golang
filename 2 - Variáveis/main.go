package main

import "fmt"

func main() {
	// Forma 1: Declaração completa (tipo explícito)
	var nome string = "João"
	var idade int = 30

	// Forma 2: Tipo inferido (Go descobre o tipo)
	var cidade = "São Paulo" // string
	var salario = 2500.50    // float64

	// Forma 3: Declaração curta (dentro de funções)
	pais := "Brasil" // string
	codigo := 55     // int

	// Forma 4: Múltiplas declarações
	var (
		dia = 25
		mes = 12
		ano = 2023
	)

	// Forma 5: Declaração curta múltipla
	altura, peso := 1.75, 70.5

	// Exibindo os valores
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Cidade:", cidade)
	fmt.Println("Salário:", salario)
	fmt.Println("País:", pais)
	fmt.Println("Código:", codigo)
	fmt.Println("Data:", dia, "/", mes, "/", ano)
	fmt.Println("Altura:", altura, "Peso:", peso)

	// Valores zero (valores padrão)
	var a int     // 0
	var b float64 // 0.0
	var c string  // ""
	var d bool    // false

	fmt.Println("\nValores zero:")
	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)

	// Constantes (valores que não mudam)
	const PI = 3.14159
	const VERSAO = "1.0.0"

	fmt.Println("\nConstantes:")
	fmt.Println("PI:", PI)
	fmt.Println("Versão:", VERSAO)
}
