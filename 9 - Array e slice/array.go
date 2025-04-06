package main

import "fmt"

func main() {
	/*
	   ARRAYS EM GO

	   Características principais:
	   - Tamanho fixo (definido na declaração)
	   - Tipo inclui o tamanho (ex: [5]int é diferente de [10]int)
	   - Por valor (cópias são independentes)
	*/

	// 1. DECLARAÇÃO BÁSICA
	var numeros [5]int // Array de 5 inteiros (todos inicializados com 0)

	// Atribuindo valores
	numeros[0] = 10
	numeros[1] = 20
	numeros[4] = 50 // Último elemento (índice 4 para array de tamanho 5)

	fmt.Println("Array numeros:", numeros) // [10 20 0 0 50]

	// 2. DECLARAÇÃO COM INICIALIZAÇÃO
	nomes := [3]string{"João", "Maria", "Carlos"}
	fmt.Println("Array nomes:", nomes) // [João Maria Carlos]

	// 3. COMPRIMENTO DO ARRAY (len)
	fmt.Println("Tamanho do array nomes:", len(nomes)) // 3

	// 4. ITERAÇÃO SOBRE ARRAY
	fmt.Println("\nIterando com for tradicional:")
	for i := 0; i < len(numeros); i++ {
		fmt.Printf("numeros[%d] = %d\n", i, numeros[i])
	}

	fmt.Println("\nIterando com range:")
	for indice, valor := range nomes {
		fmt.Printf("nomes[%d] = %s\n", indice, valor)
	}

	// 5. ARRAYS SÃO POR VALOR (CÓPIAS SÃO INDEPENDENTES)
	copia := numeros
	copia[0] = 100 // Modifica apenas a cópia

	fmt.Println("\nOriginal:", numeros) // [10 20 0 0 50]
	fmt.Println("Cópia:", copia)        // [100 20 0 0 50]

	// 6. ARRAY MULTIDIMENSIONAL
	var matriz [2][3]int
	matriz[0][0] = 1
	matriz[1][2] = 6

	fmt.Println("\nMatriz 2x3:")
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matriz[i][j])
		}
		fmt.Println()
	}

	// 7. COMPARAÇÃO DE ARRAYS
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	fmt.Println("\na == b:", a == b) // true
	fmt.Println("a == c:", a == c)   // false

	/*
	   OBSERVAÇÕES IMPORTANTES:
	   1. Arrays em Go são pouco usados diretamente (slices são mais comuns)
	   2. O tamanho é parte do tipo ([5]int ≠ [10]int)
	   3. Passagem por valor pode ser ineficiente para arrays grandes
	*/
}
