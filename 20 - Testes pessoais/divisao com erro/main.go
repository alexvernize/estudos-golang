package main

import (
	"errors"
	"fmt"
)

// Função simples que divide dois números, não retorna divisão por zero
func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisão por zero não permitida")
	}
	return a / b, nil
}

func main() {
	//Cria duas variáveis para a função divide
	divisao, erro := divisao(15, 0)
	if erro != nil {
		fmt.Println("Erro: ", erro)
	} else {
		fmt.Println("Resultado: ", divisao)
	}
}
