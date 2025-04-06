package main

import "fmt"

/*
CLOSURES EM GO

Uma closure é uma função que captura variáveis do escopo onde foi definida.
Em Go, funções anônimas podem acessar e modificar variáveis definidas fora
de seu próprio corpo, mantendo o estado entre chamadas.
*/

// contador retorna uma função closure que incrementa um contador interno
func contador() func() int {
	// i é capturado pela closure abaixo
	i := 0

	// A função anônima mantém acesso à variável i mesmo após contador() retornar
	return func() int {
		i++ // Modifica a variável capturada
		return i
	}
}

/*
ACUMULADOR COM CLOSURE

Este exemplo mostra como closures podem ser usadas para criar funções com estado.
A função acumulador mantém uma variável interna 'soma' que persiste entre chamadas.
*/
func acumulador(valorInicial int) func(int) int {
	soma := valorInicial // Variável capturada pela closure

	return func(valor int) int {
		soma += valor // Modifica o estado acumulado
		return soma
	}
}

func main() {
	/*
	   USANDO CLOSURES

	   Cada chamada a contador() cria um novo closure com seu próprio estado.
	   O estado (variável i) é isolado para cada instância.
	*/
	proximo := contador()
	fmt.Println(proximo()) // 1 (i inicializado como 0, incrementado para 1)
	fmt.Println(proximo()) // 2 (i agora vale 1, incrementado para 2)
	fmt.Println(proximo()) // 3 (continua incrementando o mesmo i)

	/*
	   CLOSURES INDEPENDENTES

	   Novo contador cria um novo closure com novo estado.
	   O contador original mantém seu próprio estado separado.
	*/
	novoContador := contador()
	fmt.Println("Novo contador:", novoContador()) // 1 (i independente)
	fmt.Println("Contador original:", proximo())  // 4 (continua do estado anterior)

	/*
	   ACUMULADOR EM AÇÃO

	   O closure mantém o estado da soma entre chamadas.
	*/
	meuAcumulador := acumulador(10)
	fmt.Println(meuAcumulador(5))  // 15 (10 + 5)
	fmt.Println(meuAcumulador(10)) // 25 (15 + 10)
	fmt.Println(meuAcumulador(20)) // 45 (25 + 20)
}
