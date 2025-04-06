package main

import "fmt"

func main() {
	/*
	   ESTRUTURAS DE CONTROLE IF/ELSE EM GO

	   O if/else permite controlar o fluxo do programa com base em condições.
	   Em Go, a sintaxe é simples e não usa parênteses nas condições.
	*/

	// 1. IF BÁSICO
	idade := 18
	fmt.Println("1. Verificação básica de idade:")

	if idade >= 18 {
		fmt.Println("Você é maior de idade!")
	}

	// 2. IF COM ELSE
	temperatura := 25.5
	fmt.Println("\n2. Verificação de temperatura:")

	if temperatura > 30 {
		fmt.Println("Está muito quente!")
	} else {
		fmt.Println("A temperatura está agradável.")
	}

	// 3. IF COM ELSE IF (MÚLTIPLAS CONDIÇÕES)
	nota := 7.5
	fmt.Println("\n3. Avaliação de nota:")

	if nota >= 9 {
		fmt.Println("Excelente!")
	} else if nota >= 7 {
		fmt.Println("Bom!")
	} else if nota >= 5 {
		fmt.Println("Recuperação")
	} else {
		fmt.Println("Reprovado")
	}

	// 4. IF COM INICIALIZAÇÃO (if init)
	// Podemos declarar variáveis antes da condição
	fmt.Println("\n4. If com inicialização:")

	if saldo := 100.0; saldo < 0 {
		fmt.Println("Saldo negativo!")
	} else if saldo < 50 {
		fmt.Println("Saldo baixo")
	} else {
		fmt.Printf("Saldo positivo: R$%.2f\n", saldo)
	}

	// A variável saldo só existe dentro do bloco if/else
	// fmt.Println(saldo) // Erro: saldo não definido

	// 5. OPERADORES LÓGICOS (&&, ||, !)
	hora := 14
	fmt.Println("\n5. Operadores lógicos:")

	if hora >= 6 && hora < 12 {
		fmt.Println("Bom dia!")
	} else if hora >= 12 && hora < 18 {
		fmt.Println("Boa tarde!")
	} else {
		fmt.Println("Boa noite!")
	}

	// 6. IF COM FUNÇÕES
	fmt.Println("\n6. If com funções:")

	if resultado := divisao(10, 2); resultado > 0 {
		fmt.Println("Resultado positivo:", resultado)
	} else {
		fmt.Println("Resultado negativo ou zero")
	}

	// 7. IF PARA TRATAR ERROS (padrão comum em Go)
	fmt.Println("\n7. Tratamento de erros:")

	if resultado, err := divisao(10, 0); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

// Função auxiliar para exemplos
func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("não pode dividir por zero")
	}
	return a / b, nil
}
