package main

import "fmt"

func main() {
	/*
	   LAÇOS (LOOPS) EM GO

	   Go tem apenas uma estrutura de loop: o 'for', mas ele pode ser usado
	   de várias maneiras equivalentes aos loops 'while' e 'do-while' de outras linguagens.
	*/

	// 1. FOR BÁSICO (estilo C)
	fmt.Println("1. For básico (estilo C):")
	for i := 0; i < 5; i++ {
		fmt.Println("Contador:", i)
	}

	// 2. FOR COMO WHILE (somente condição)
	fmt.Println("\n2. For como while:")
	j := 0
	for j < 3 {
		fmt.Println("While-style counter:", j)
		j++
	}

	// 3. LOOP INFINITO (com break)
	fmt.Println("\n3. Loop infinito (com break):")
	k := 0
	for {
		if k >= 3 {
			break // Sai do loop
		}
		fmt.Println("Loop infinito counter:", k)
		k++
	}

	// 4. CONTINUE (pula iteração)
	fmt.Println("\n4. Usando continue:")
	for n := 0; n < 5; n++ {
		if n%2 == 0 { // Pula números pares
			continue
		}
		fmt.Println("Número ímpar:", n)
	}

	// 5. RANGE (iterar sobre coleções)
	fmt.Println("\n5. Range com slices:")
	frutas := []string{"maçã", "banana", "laranja"}
	for index, fruta := range frutas {
		fmt.Printf("Index: %d, Fruta: %s\n", index, fruta)
	}

	// Range ignorando o índice
	fmt.Println("\nRange ignorando índice:")
	for _, fruta := range frutas {
		fmt.Println("Fruta:", fruta)
	}

	// Range com maps
	fmt.Println("\nRange com maps:")
	capitais := map[string]string{
		"Brasil": "Brasília",
		"EUA":    "Washington",
		"Japão":  "Tóquio",
	}
	for pais, capital := range capitais {
		fmt.Printf("A capital do %s é %s\n", pais, capital)
	}

	/*
	   DICAS IMPORTANTES:
	   1. Não existe while ou do-while em Go - use 'for' no lugar
	   2. O range retorna dois valores: índice/chave e valor
	   3. Use _ para ignorar um dos valores retornados
	   4. A ordem de iteração em maps não é garantida
	*/
}
