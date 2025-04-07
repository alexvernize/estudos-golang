// Concorrência é diferente de paralelismo
// Concorrência é quando várias tarefas estão progredindo ao mesmo tempo
// Paralelismo é quando várias tarefas estão sendo executadas ao mesmo tempo
// Em Go, a concorrência é implementada através de goroutines e canais
// Goroutines são funções que podem ser executadas concorrentemente
// Elas são leves e podem ser criadas facilmente
// Canais são usados para comunicação entre goroutines
// Eles permitem que você envie e receba dados entre goroutines
// Canais são seguros para uso concorrente e garantem que os dados sejam enviados e recebidos na ordem correta

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
	// Criando goroutines
	go func() {
		println("Goroutine 1")
	}()
	go func() {
		println("Goroutine 2")
	}()
	go func() {
		println("Goroutine 3")
	}()

	// Esperando as goroutines terminarem
	var input string
	fmt.Scanln(&input)

	escrever("Olá, mundo!")
}
