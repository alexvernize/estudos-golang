package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
func main() {
	// Um waitgroup é um contador que espera até que todas as goroutines terminem
	// Ele é usado para sincronizar a execução de várias goroutines
	// Ele é útil quando você precisa esperar que várias goroutines terminem antes de continuar
	// a execução do programa
	// Ele é usado para evitar que o programa termine antes que todas as goroutines terminem
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Adicionando duas goroutines

	// Criando uma goroutine anônima
	// A goroutine anônima é uma função que é executada em paralelo com o restante do programa
	// Ela é útil quando você precisa executar uma função em paralelo, mas não precisa
	// de um nome para a função
	go func() {
		defer waitGroup.Done() // Defer é usado para garantir que o contador seja decrementado
		escrever("Olá, mundo!")
	}()
	go func() {
		defer waitGroup.Done() // Defer é usado para garantir que o contador seja decrementado
		escrever("Programando em Golang!")
	}()
	waitGroup.Wait() // Esperando todas as goroutines terminarem
}
