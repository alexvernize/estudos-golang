package main

import (
	"errors"
	"fmt"
)

// Função simples que recebe e divide dois números. Esta função retorna dois valores,
// o valor da divisão e nil, caso o divisor seja zero, o número retornado será zero e um erro
// também é será retornado.
func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisão por zero não permitida")
	}
	return a / b, nil
}

func main() {
	//Cria duas variáveis para a função divisao
	divisao, erro := divisao(15, 0)
	if erro != nil {
		fmt.Println("Erro: ", erro)
	} else {
		fmt.Println("Resultado: ", divisao)
	}
}
