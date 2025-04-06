package main

import "fmt"

func main() {
	/*
	   SLICES EM GO

	   Características principais:
	   - Tamanho dinâmico (pode crescer ou diminuir)
	   - Referência a um array subjacente
	   - Tipo não inclui tamanho ([]int)
	   - Muito mais utilizados que arrays
	*/

	// 1. DECLARAÇÃO BÁSICA
	var slice1 []int                                        // slice vazio (nil)
	fmt.Println("slice1:", slice1, "é nil?", slice1 == nil) // [] true

	// 2. INICIALIZAÇÃO COM VALORES
	slice2 := []string{"Go", "Python", "JavaScript"}
	fmt.Println("\nslice2:", slice2) // [Go Python JavaScript]

	// 3. CRIANDO A PARTIR DE ARRAY
	array := [5]int{10, 20, 30, 40, 50}
	slice3 := array[1:4]                        // Do índice 1 ao 3 (4 não incluso)
	fmt.Println("\nslice3 (do array):", slice3) // [20 30 40]

	// 4. COMPRIMENTO (len) E CAPACIDADE (cap)
	fmt.Println("Tamanho (len):", len(slice3))    // 3
	fmt.Println("Capacidade (cap):", cap(slice3)) // 4 (do índice 1 até o fim do array)

	// 5. MODIFICANDO O SLICE (AFETA O ARRAY ORIGINAL)
	slice3[0] = 200
	fmt.Println("\nArray modificado:", array) // [10 200 30 40 50]
	fmt.Println("Slice modificado:", slice3)  // [200 30 40]

	// 6. FUNÇÃO MAKE (CRIA SLICE COM TAMANHO E CAPACIDADE)
	slice4 := make([]int, 3, 5) // Tamanho 3, capacidade 5
	fmt.Println("\nslice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))

	// 7. ADICIONANDO ELEMENTOS (APPEND)
	slice4 = append(slice4, 10)
	slice4 = append(slice4, 20)
	slice4 = append(slice4, 30) // Aumenta a capacidade automaticamente
	fmt.Println("\nApós append:", slice4, "len:", len(slice4), "cap:", cap(slice4))

	// 8. COPIANDO SLICES (COPY)
	src := []int{1, 2, 3}
	dest := make([]int, 2)
	n := copy(dest, src)
	fmt.Println("\nCopiados", n, "elementos:", dest) // [1 2]

	// 9. SLICE INTERNO (NÃO AFETA O ORIGINAL)
	slice5 := []int{100, 200, 300, 400, 500}
	slice6 := slice5[:3] // Novo slice referenciando os mesmos dados
	slice6[0] = 999      // Modifica ambos os slices

	// Para criar um slice independente, use copy ou append em um slice vazio
	slice7 := append([]int{}, slice5...)
	slice7[0] = 111 // Só modifica slice7

	fmt.Println("\nslice5:", slice5) // [999 200 300 400 500]
	fmt.Println("slice6:", slice6)   // [999 200 300]
	fmt.Println("slice7:", slice7)   // [111 200 300 400 500]

	// 10. SLICES MULTIDIMENSIONAIS
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("\nMatriz como slice de slices:")
	for _, linha := range matriz {
		fmt.Println(linha)
	}

	/*
	   OBSERVAÇÕES IMPORTANTES:
	   1. Slices são referências (cuidado com compartilhamento não intencional)
	   2. Append pode realocar memória se a capacidade for excedida
	   3. Use make quando souber o tamanho aproximado necessário
	   4. copy é útil para evitar compartilhamento de dados
	*/
}
