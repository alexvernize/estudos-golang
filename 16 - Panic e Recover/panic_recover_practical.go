package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
	   EXEMPLO PRÁTICO: DIVISÃO POR ZERO
	   - Se tentarmos dividir por zero, causamos um panic.
	   - Usamos recover para tratar o erro e evitar que o programa quebre.
	*/

	resultado, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	// Testando um caso válido
	resultado, err = dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func dividir(a, b int) (resultado int, err error) {
	/*
	   Função que trata divisão por zero usando panic/recover.
	*/
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("não é possível dividir por zero")
		}
	}()

	if b == 0 {
		panic("divisão por zero") // Causa panic se b == 0
	}

	return a / b, nil // Retorna normalmente se não houver erro
}
