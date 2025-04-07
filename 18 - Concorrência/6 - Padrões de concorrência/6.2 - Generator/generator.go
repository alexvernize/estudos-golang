package main

import (
	"fmt"
	"time"
)

// O padrão generator é um padrão de design que permite gerar uma sequência de valores
// sob demanda, em vez de calcular todos os valores de uma vez.
// Isso é útil quando você não sabe quantos valores precisará ou quando os valores são caros para calcular.
// O padrão generator é semelhante ao padrão iterator, mas o generator é mais flexível
// e pode ser usado em situações onde o iterator não é adequado.
// O padrão generator é uma função que retorna um canal (channel) e pode ser usado para gerar
// uma sequência de valores. O canal é usado para enviar os valores gerados para o código que
// está chamando o generator. O código que chama o generator pode usar um loop para receber os
// valores do canal e processá-los um por um. Isso permite que o código que chama o generator
// processe os valores à medida que eles são gerados, em vez de esperar que todos os valores sejam
// gerados de uma vez.
func escrever(texto string) <-chan string {
	canal := make(chan string) // Criando um canal de string
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto) // Enviando mensagens para o canal
			time.Sleep(time.Millisecond * 500)                // Esperando 500ms para simular um atraso
		}
		// for i := 0; i < 5; i++ {
		// 	canal <- texto
		// }
		// close(canal) // Fechando o canal
	}()
	return canal // Retornando o canal
}

func main() {
	canal := escrever("Olá, mundo!") // Chamando a função escrever e recebendo o canal
	// O código que chama o generator pode usar um loop para receber os valores do canal e processá-los um por um.
	// Isso permite que o código que chama o generator processe os valores à medida que eles são gerados, em vez de esperar que todos os valores sejam gerados de uma vez.
	for i := 0; i < 10; i++ {
		mensagem := <-canal // Recebendo a mensagem do canal
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do programa!")
}
