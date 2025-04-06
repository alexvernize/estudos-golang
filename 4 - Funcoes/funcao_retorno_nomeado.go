package main

import (
	"fmt"
	"math"
)

/*
RETORNOS NOMEADOS EM GO

Permitem definir nomes para os valores de retorno na assinatura da função.
Benefícios:
- Melhor documentação
- Código mais legível
- Valores são inicializados com zero value
- Pode usar return simples (sem especificar valores)
*/

// 1. Retorno nomeado simples
func calculaArea(raio float64) (area float64) {
	// area é automaticamente inicializada com 0.0
	area = math.Pi * raio * raio
	return // Não precisa especificar o valor (usa o nomeado)
}

// 2. Múltiplos retornos nomeados
func divide(dividendo, divisor float64) (resultado float64, erro error) {
	if divisor == 0 {
		erro = fmt.Errorf("não pode dividir por zero")
		return // Retorna 0.0 e o erro
	}
	resultado = dividendo / divisor
	return // Retorna o resultado e nil
}

// 3. Retorno nomeado em defer
func exemploDefer(valor int) (x int) {
	defer func() {
		x *= 2 // Modifica o retorno nomeado
	}()
	x = valor + 5
	return // x será (valor + 5) * 2
}

func main() {
	// 1. Usando retorno nomeado simples
	fmt.Println("Área do círculo:", calculaArea(5)) // 78.53981633974483

	// 2. Usando múltiplos retornos nomeados
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", res) // 5
	}

	// Tentando dividir por zero
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Erro:", err) // Exibe o erro
	}

	// 3. Demonstração com defer
	fmt.Println("Exemplo defer:", exemploDefer(10)) // 30 ((10 + 5) * 2)

	/*
	   QUANDO USAR RETORNOS NOMEADOS:
	   - Quando os nomes adicionam clareza
	   - Para documentar o significado dos retornos
	   - Quando usando defer para modificar retornos
	   - Em funções com múltiplos retornos do mesmo tipo
	*/
}
