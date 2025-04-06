package main

import "fmt"

/*
FUNÇÕES VARIÁDICAS

Funções variádicas aceitam um número variável de argumentos do mesmo tipo.
Em Go, isso é indicado com ... antes do tipo do parâmetro.
*/

// somaVariadica pode receber zero ou mais inteiros
func somaVariadica(numeros ...int) int {
	total := 0

	// numeros é tratado como um slice dentro da função
	for _, num := range numeros {
		total += num
	}

	return total
}

/*
FUNÇÃO COM PARÂMETROS NORMAIS E VARIÁDICOS

Quando combinados, o parâmetro variádico deve ser o último.
*/
func imprimirDados(prefixo string, valores ...float64) {
	fmt.Print(prefixo, ": ")
	for _, v := range valores {
		fmt.Printf("%.2f ", v)
	}
	fmt.Println()
}

func main() {
	// CHAMANDO FUNÇÃO VARIÁDICA
	fmt.Println("Soma de 1 a 4:", somaVariadica(1, 2, 3, 4)) // 10
	fmt.Println("Soma sem argumentos:", somaVariadica())     // 0

	// PASSANDO SLICE PARA FUNÇÃO VARIÁDICA
	nums := []int{10, 20, 30}
	fmt.Println("Soma do slice:", somaVariadica(nums...)) // 60

	// USANDO PARÂMETROS NORMAIS E VARIÁDICOS
	imprimirDados("Valores", 1.5, 2.7, 3.8, 4.2)

	/*
	   CASOS DE USO COMUNS:
	   1. Funções de soma/média
	   2. Funções de formatação/impressão
	   3. Funções utilitárias que processam listas
	*/
}
