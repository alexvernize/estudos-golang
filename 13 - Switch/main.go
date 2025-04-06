package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		ESTRUTURA SWITCH EM GO

		O switch é uma forma mais limpa de escrever múltiplas condições if/else.
		Em Go, o switch tem algumas particularidades interessantes.
	*/

	// 1. SWITCH BÁSICO (COM NÚMEROS)
	diaDaSemana := 3
	fmt.Println("1. Switch básico:")
	switch diaDaSemana {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Dia inválido")
	}

	// 2. SWITCH COM EXPRESSÕES
	numero := 15
	fmt.Println("\n2. Switch com expressões:")
	switch {
	case numero < 10:
		fmt.Println("Número pequeno")
	case numero >= 10 && numero < 20:
		fmt.Println("Número médio")
	default:
		fmt.Println("Número grande")
	}

	// 3. SWITCH COM MULTIPLOS VALORES
	letra := "b"
	fmt.Println("\n3. Switch com múltiplos valores:")
	switch letra {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vogal")
	case "b", "c", "d", "f", "g":
		fmt.Println("Consoante")
	default:
		fmt.Println("Não é uma letra válida")
	}

	// 4. SWITCH COM INICIALIZAÇÃO
	fmt.Println("\n4. Switch com inicialização:")
	switch hora := time.Now().Hour(); {
	case hora < 12:
		fmt.Println("Bom dia!")
	case hora < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}

	// 5. SWITCH COM TIPOS (TYPE SWITCH)
	fmt.Println("\n5. Switch com tipos:")
	qualquerCoisa := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("É um inteiro")
		case float64:
			fmt.Println("É um float")
		case string:
			fmt.Println("É uma string")
		default:
			fmt.Printf("Tipo desconhecido: %T\n", t)
		}
	}

	qualquerCoisa(10)
	qualquerCoisa(3.14)
	qualquerCoisa("texto")
	qualquerCoisa(true)

	/*
		DIFERENÇAS IMPORTANTES:
		1. Não precisa do 'break' - Go sai automaticamente do switch
		2. Pode usar vírgula para múltiplos valores
		3. O 'default' é opcional
		4. Pode usar sem expressão (como if/else)
		5. Pode conter inicialização
	*/
}
