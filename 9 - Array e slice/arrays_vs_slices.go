package main

import "fmt"

func main() {
	/*
	   COMPARAÇÃO ENTRE ARRAYS E SLICES

	   Este exemplo mostra as diferenças práticas entre arrays e slices
	*/

	// 1. DECLARAÇÃO E INICIALIZAÇÃO
	array := [3]int{1, 2, 3} // Array com tamanho fixo
	slice := []int{1, 2, 3}  // Slice com tamanho dinâmico

	fmt.Println("Array:", array) // [1 2 3]
	fmt.Println("Slice:", slice) // [1 2 3]

	// 2. TAMANHO E TIPO
	fmt.Printf("\nTipo do array: %T\n", array) // [3]int
	fmt.Printf("Tipo do slice: %T\n", slice)   // []int

	// 3. PASSAGEM PARA FUNÇÕES
	modificaArray(array) // Passagem por valor (cópia)
	modificaSlice(slice) // Passagem por referência

	fmt.Println("\nApós modificação:")
	fmt.Println("Array (não mudou):", array)  // [1 2 3]
	fmt.Println("Slice (modificado):", slice) // [100 2 3]

	// 4. CAPACIDADE E FLEXIBILIDADE
	// array = append(array, 4) // Erro: arrays não suportam append
	slice = append(slice, 4, 5, 6)             // Slices podem crescer
	fmt.Println("\nSlice após append:", slice) // [100 2 3 4 5 6]

	// 5. RELAÇÃO ENTRE ARRAYS E SLICES
	novoArray := [5]string{"A", "B", "C", "D", "E"}
	novoSlice := novoArray[1:4] // Slice referenciando parte do array

	fmt.Println("\nArray original:", novoArray) // [A B C D E]
	fmt.Println("Slice do array:", novoSlice)   // [B C D]

	novoSlice[0] = "Z" // Modifica tanto o slice quanto o array
	fmt.Println("\nApós modificação via slice:")
	fmt.Println("Array:", novoArray) // [A Z C D E]
	fmt.Println("Slice:", novoSlice) // [Z C D]

	/*
	   QUANDO USAR CADA UM:
	   - Arrays: Quando você precisa de tamanho fixo e imutável
	   - Slices: Em 99% dos casos, quando precisa de flexibilidade
	*/
}

func modificaArray(a [3]int) {
	a[0] = 100 // Modifica apenas a cópia local
}

func modificaSlice(s []int) {
	s[0] = 100 // Modifica o slice original
}
