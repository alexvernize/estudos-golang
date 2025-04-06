package main

import (
	"fmt"
	"os"
)

/*
FUNÇÃO INIT EM GO

A função init é uma função especial que:
- É executada automaticamente antes da main()
- Pode haver uma por arquivo
- É executada na ordem de importação dos pacotes
- Usada para inicializações globais
*/

// 1. init simples
func init() {
	fmt.Println("1. Inicializando o programa...")
}

// 2. init com verificação de ambiente
func init() {
	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Modo debug ativado")
	}
}

var config string

// 3. init que inicializa variáveis globais
func init() {
	config = "configuração carregada"
	fmt.Println("Configuração inicializada:", config)
}

func main() {
	fmt.Println("Função main executando...")
	fmt.Println("Usando configuração:", config)

	/*
	   ORDEM DE EXECUÇÃO:
	   1. Todas as funções init do pacote main
	   2. Função main()

	   PARA PACOTES IMPORTADOS:
	   1. Todas as funções init dos pacotes importados
	   2. Funções init do pacote atual
	   3. Função main()
	*/
}

// 4. Outra init (executada na ordem de declaração)
func init() {
	fmt.Println("2. Continuando inicialização...")
}

/*
CASOS DE USO COMUNS PARA INIT:
- Carregar configurações
- Verificar dependências
- Inicializar conexões
- Registrar componentes
- Validar ambiente
*/
