package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	// Como o canal espera uma mensagem chegar ele não vai continuar a execução do loop
	// e na primeira mensagem que ele receber ele vai imprimir e terminar o código
	//time.Sleep(time.Second * 5) // Esperando 2 segundos para simular um atraso
	for i := 0; i < 5; i++ {
		//Para passar um valor para o canal, você usa o operador de atribuição "<-"
		canal <- texto
		time.Sleep(time.Second)

		//Para resolver possíveis deadlocks, você pode usar o operador "close" para fechar o canal
		//depois que você terminar de enviar mensagens para o canal
	}
	close(canal) // Fechando o canal
}
func main() {
	//Channel é um canal de comunicação entre goroutines
	// Ele é usado para enviar e receber dados entre goroutines
	// Ele é seguro para uso concorrente e garante que os dados sejam enviados e recebidos na ordem correta
	// Ele é usado para evitar condições de corrida e garantir que os dados sejam consistentes
	// Ele é usado para sincronizar a execução de várias goroutines
	// Ele é usado para evitar que o programa termine antes que todas as goroutines terminem
	// Ele é usado para evitar que o programa fique bloqueado esperando por uma goroutine
	// Ele é usado para evitar que o programa fique bloqueado esperando por um canal

	canal := make(chan string) // Criando um canal de string
	go escrever("Olá, mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	//Dessa forma ele espera todo o loop da função escrever terminar
	//porém ele gera um deadlock. O deadlock acontece quando você não tem mais nenhum lugar enviando mensagens para o canal
	//mas o canal ainda está esperando receber mensagens
	// for {
	// 	mensagem, aberto := <-canal // Para receber um valor do canal, você usa o operador de atribuição "<-"
	// 	if !aberto {                // Se o canal estiver fechado, ele não vai entrar no loop
	// 		fmt.Println("Canal fechado")
	// 		break
	// 	}
	// 	fmt.Println(mensagem) // Imprimindo a mensagem recebida do canal
	// }
	//Uma forma mais simples de fazer isso é usar o operador "for range" para iterar sobre o canal
	for mensagem := range canal { // Para receber um valor do canal, você usa o operador de atribuição "<-"
		fmt.Println(mensagem) // Imprimindo a mensagem recebida do canal
	}
	fmt.Println("Fim do programa!")

	// mensagem := <-canal // Para receber um valor do canal, você usa o operador de atribuição "<-"
	// println(mensagem)   // Imprimindo a mensagem recebida do canal
}
