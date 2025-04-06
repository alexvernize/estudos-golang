package main

import "fmt"

func main() {
	/*
	   MAPS EM GO

	   Maps são estruturas de dados que armazenam pares chave-valor.
	   - Similar a dicionários em Python ou objetos em JavaScript
	   - Chaves devem ser de tipos comparáveis (==, !=)
	   - Valores podem ser de qualquer tipo
	   - Referências (como slices)
	*/

	// 1. DECLARAÇÃO E INICIALIZAÇÃO
	var capitais map[string]string // Declaração (map nil)

	// capitais["Brasil"] = "Brasília" // PANIC: map não inicializado

	// Inicialização com make
	capitais = make(map[string]string)
	capitais["Brasil"] = "Brasília" // Agora funciona

	// Declaração e inicialização direta
	populacao := map[string]int{
		"São Paulo":      12300000,
		"Rio de Janeiro": 6748000,
	}

	fmt.Println("Map de capitais:", capitais)
	fmt.Println("Map de população:", populacao)

	// 2. OPERAÇÕES BÁSICAS
	// Adicionando/Atualizando
	populacao["Belo Horizonte"] = 2531000

	// Acessando valores
	sp := populacao["São Paulo"]
	fmt.Println("\nPopulação de SP:", sp)

	// Verificando existência
	if capital, existe := capitais["Brasil"]; existe {
		fmt.Println("Capital do Brasil:", capital)
	}

	// Removendo itens
	delete(populacao, "Rio de Janeiro")
	fmt.Println("Após remover Rio:", populacao)

	// 3. ITERAÇÃO EM MAPS
	fmt.Println("\nIterando sobre o map de capitais:")
	for pais, capital := range capitais {
		fmt.Printf("%s -> %s\n", pais, capital)
	}

	// 4. MAP ZERO VALUE (nil)
	var mapNil map[int]string
	fmt.Println("\nMap nil:", mapNil)
	// mapNil[1] = "um" // PANIC: assignment to nil map

	// 5. COMPRIMENTO DO MAP
	fmt.Println("Número de capitais:", len(capitais)) // 1

	/*
	   OBSERVAÇÕES IMPORTANTES:
	   1. Maps são referências (como slices)
	   2. A ordem de iteração não é garantida
	   3. Operações são concorrentemente inseguras
	*/
}
