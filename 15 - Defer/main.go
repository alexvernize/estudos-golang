package main

import "fmt"

func main() {
	/*
	   DEFER EM GO

	   O 'defer' adia a execução de uma função até o final da função atual.
	   É muito usado para:
	   - Liberar recursos (arquivos, conexões)
	   - Fechar locks
	   - Finalizar transações
	   - Garantir limpeza
	*/

	// 1. DEFER BÁSICO
	fmt.Println("\n1. Ordem com defer:")
	fmt.Println("Primeira mensagem")
	defer fmt.Println("Mensagem deferida") // Será executada por último
	fmt.Println("Segunda mensagem")

	// 2. DEFER EM LOOP (cada defer é empilhado)
	fmt.Println("\n2. Defer em loop:")
	for i := 0; i < 3; i++ {
		defer fmt.Println("Defer no loop:", i) // LIFO (último a entrar, primeiro a sair)
	}
	fmt.Println("Loop concluído")

	// 3. DEFER COM FUNÇÕES
	fmt.Println("\n3. Defer com funções:")
	defer func() {
		fmt.Println("Função anônima deferida")
	}()
	fmt.Println("Função principal executando")

	// 4. DEFER COM PARÂMETROS (os valores são capturados no momento do defer)
	fmt.Println("\n4. Defer com parâmetros:")
	a := 10
	defer fmt.Println("Valor de a no defer:", a) // Captura o valor atual (10)
	a = 20
	fmt.Println("Valor de a modificado:", a)

	// 5. USO PRÁTICO - SIMULANDO ABERTURA DE ARQUIVO
	fmt.Println("\n5. Exemplo prático (simulação):")
	abrirArquivo()
	fmt.Println("Operação concluída")

	/*
	   REGRAS IMPORTANTES:
	   1. Os defers são executados na ordem inversa (LIFO)
	   2. Os argumentos são avaliados no momento do defer
	   3. Defers são executados mesmo que haja panic
	   4. Útil para limpeza de recursos
	*/
}

func abrirArquivo() {
	fmt.Println("Abrindo arquivo...")
	defer fmt.Println("Fechando arquivo...") // Garante que será executado no final

	fmt.Println("Escrevendo no arquivo...")
	// Simulando erro
	if true { // Altere para false para ver o fluxo sem erro
		panic("Erro durante a escrita") // Mesmo com panic, o defer executa
	}
	fmt.Println("Concluído com sucesso") // Não será executado em caso de panic
}
