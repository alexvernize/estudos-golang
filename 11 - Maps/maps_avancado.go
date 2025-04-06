package main

import "fmt"

func main() {
	/*
	   USOS AVANÇADOS DE MAPS
	*/

	// 1. MAP DE MAPS (TABELA)
	cidades := map[string]map[string]float64{
		"Brasil": {
			"São Paulo":      12300000,
			"Rio de Janeiro": 6748000,
		},
		"EUA": {
			"Nova York":   8419000,
			"Los Angeles": 3976000,
		},
	}

	fmt.Println("População de NY:", cidades["EUA"]["Nova York"])

	// 2. MAP COM CHAVE COMPOSTA
	type Coordenada struct {
		Lat, Long float64
	}
	locais := map[Coordenada]string{
		{Lat: -23.5505, Long: -46.6333}: "São Paulo",
		{Lat: -22.9068, Long: -43.1729}: "Rio de Janeiro",
	}
	fmt.Println("\nLocal em (-23.5505, -46.6333):", locais[Coordenada{-23.5505, -46.6333}])

	// 3. MAP DE SLICES
	timesPorPais := map[string][]string{
		"Brasil":  {"Flamengo", "Palmeiras", "Corinthians"},
		"Espanha": {"Real Madrid", "Barcelona", "Atlético de Madrid"},
	}
	fmt.Println("\nTimes do Brasil:", timesPorPais["Brasil"])

	// Adicionando a um slice dentro do map
	timesPorPais["Brasil"] = append(timesPorPais["Brasil"], "São Paulo")
	fmt.Println("Times atualizados:", timesPorPais["Brasil"])

	// 4. MAP COMO CACHE SIMPLES
	cache := make(map[string]int)
	palavras := []string{"gato", "cachorro", "gato", "pássaro", "cachorro"}

	fmt.Println("\nContagem de palavras:")
	for _, palavra := range palavras {
		cache[palavra]++
	}
	fmt.Println(cache) // map[cachorro:2 gato:2 pássaro:1]

	// 5. MAP COMO SET (CONJUNTO)
	set := make(map[string]bool)
	elementos := []string{"A", "B", "A", "C", "B"}

	for _, elem := range elementos {
		set[elem] = true
	}
	fmt.Println("\nElementos únicos:")
	for elem := range set {
		fmt.Println(elem) // A, B, C (ordem não garantida)
	}

	// Verificando existência no set
	fmt.Println("Contém 'A'?", set["A"]) // true
	fmt.Println("Contém 'D'?", set["D"]) // false
}
