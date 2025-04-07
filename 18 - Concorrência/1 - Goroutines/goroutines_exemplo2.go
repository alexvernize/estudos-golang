/*
Esse código abaixo mostra a chamada de uma função que imprime uma string a cada segundo
A função é chamada duas vezes, mas como não há controle de concorrência, as duas chamadas
vão imprimir a string ao mesmo tempo, o que pode causar confusão
Para evitar isso, podemos usar goroutines para executar as chamadas de forma concorrente
Isso significa que as duas chamadas vão ser executadas ao mesmo tempo, mas cada uma vai imprimir
a string a cada segundo, sem interferir uma na outra
Para isso, basta adicionar a palavra-chave "go" antes da chamada da função
Isso vai fazer com que a função seja executada em uma goroutine separada
package main
import (
	"fmt"
	"time"
)
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
func main() {
	escrever("Olá, mundo!")
	escrever("Programando em Golang!")
} */

package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
func main() {
	// Criando goroutines você diz ao programa que você quer executar essa função em paralelo
	// e não esperar ela terminar para continuar a execução do programa
	// Não faz sentido colocar mais de uma goroutine para a mesma função
	// porque ela vai ficar em loop infinito
	go escrever("Olá, mundo!") //goroutine só adicionar a palavra go na frente da função
	escrever("Programando em Golang!")
}
