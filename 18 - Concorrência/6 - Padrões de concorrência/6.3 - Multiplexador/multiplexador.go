// Um multiplexador em Go é uma função que permite enviar mensagens de vários canais para um único canal.
// Isso é útil quando você tem várias goroutines que estão enviando mensagens para diferentes canais
// e você quer receber todas essas mensagens em um único canal.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem1 := <-canal1:
				canalDeSaida <- mensagem1
			case mensagem2 := <-canal2:
				canalDeSaida <- mensagem2
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}

func main() {
	canal := multiplexar(escrever("Olá, mundo!"), escrever("Programando em Golang!"))
	for i := 0; i < 10; i++ {
		mensagem := <-canal
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa!")
}
