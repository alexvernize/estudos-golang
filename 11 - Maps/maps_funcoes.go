package main

import "fmt"

func main() {
	/*
	   MAPS COM FUNÇÕES
	*/

	// 1. MAP DE FUNÇÕES (DISPATCH TABLE)
	operacoes := map[string]func(int, int) int{
		"soma":          func(a, b int) int { return a + b },
		"subtracao":     func(a, b int) int { return a - b },
		"multiplicacao": func(a, b int) int { return a * b },
	}

	resultado := operacoes["soma"](10, 5)
	fmt.Println("10 + 5 =", resultado)

	// Verificando se operação existe
	if op, existe := operacoes["divisao"]; existe {
		fmt.Println("10 / 5 =", op(10, 5))
	} else {
		fmt.Println("\nOperação divisão não existe")
	}

	// 2. PASSANDO MAPS PARA FUNÇÕES
	contagem := map[string]int{
		"a": 3,
		"b": 2,
	}
	fmt.Println("\nContagem original:", contagem)
	incrementarContagem(contagem, "a")
	incrementarContagem(contagem, "c")
	fmt.Println("Contagem atualizada:", contagem)

	// 3. RETORNANDO MAPS DE FUNÇÕES
	config := carregarConfig()
	fmt.Println("\nConfigurações:", config)
}

func incrementarContagem(m map[string]int, chave string) {
	/*
	   Modifica o map original (maps são referências)
	*/
	m[chave]++
}

func carregarConfig() map[string]string {
	/*
	   Retorna um novo map com configurações
	*/
	return map[string]string{
		"host":    "localhost",
		"porta":   "8080",
		"timeout": "30s",
	}
}
