package main // Todo programa executável em Go começa com o pacote 'main'

import (
	"fmt"

	// Importando nossos pacotes locais
	"pacotes/matematica"
	"pacotes/saudacao"
)

func main() {
	// Usando funções do pacote de saudação
	saudacao.Ola()           // Função pública (começa com letra maiúscula)
	saudacao.OlaNome("João") // Outra função pública

	// Usando funções do pacote matemática
	soma := matematica.Soma(10, 20)
	fmt.Printf("Soma: %d\n", soma)

	area := matematica.AreaRetangulo(5.0, 3.0)
	fmt.Printf("Área do retângulo: %.2f\n", area)

	// Tentar usar uma função não exportada causaria erro de compilação
	// matematica.subtracao(10, 5) // Erro: subtracao não é exportado
}
