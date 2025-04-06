package main

import "fmt"

/*
PONTEIROS EM FUNÇÕES

Ponteiros permitem que funções modifiquem valores fora de seu escopo.
Em Go, quando passamos um ponteiro para uma função, estamos passando
uma referência ao valor original, não uma cópia.
*/

// incrementoPorValor recebe uma cópia do inteiro
// Modificações não afetam o valor original
func incrementoPorValor(n int) {
	n++ // Modifica apenas a cópia local
	fmt.Println("Dentro de incrementoPorValor:", n)
}

/*
incrementoPorReferencia recebe um ponteiro para inteiro
O operador * é usado para:
1. Declarar o tipo do parâmetro (*int - ponteiro para int)
2. Acessar o valor apontado (*n - dereferenciação)
*/
func incrementoPorReferencia(n *int) {
	*n++ // Modifica o valor original
	fmt.Println("Dentro de incrementoPorReferencia:", *n)
}

/*
RETORNANDO PONTEIROS

É seguro retornar ponteiros para valores criados dentro da função.
O Go gerencia a alocação de memória automaticamente (escape analysis).
*/
func novaPessoa(nome string, idade int) *struct {
	nome  string
	idade int
} {
	// Cria uma struct anônima
	p := struct {
		nome  string
		idade int
	}{
		nome:  nome,
		idade: idade,
	}

	// Retorna o endereço da struct (ponteiro)
	return &p
}

func main() {
	// PASSAGEM POR VALOR VS REFERÊNCIA
	numero := 10

	incrementoPorValor(numero)
	fmt.Println("Após incrementoPorValor:", numero) // 10 (inalterado)

	// & obtém o endereço de memória (cria um ponteiro)
	incrementoPorReferencia(&numero)
	fmt.Println("Após incrementoPorReferencia:", numero) // 11 (modificado)

	// USANDO PONTEIROS RETORNADOS
	pessoa := novaPessoa("Maria", 30)
	fmt.Printf("%s tem %d anos\n", pessoa.nome, pessoa.idade)

	// Modificando através do ponteiro
	pessoa.idade = 31
	fmt.Printf("Agora %s tem %d anos\n", pessoa.nome, pessoa.idade)

	/*
	   POR QUE USAR PONTEIROS?
	   1. Para modificar valores originais
	   2. Para evitar cópia de grandes estruturas
	   3. Para compartilhar estado entre funções
	*/
}
