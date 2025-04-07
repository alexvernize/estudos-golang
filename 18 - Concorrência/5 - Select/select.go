package main

import "time"

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// O select é muito parecido com o switch, mas ele é usado para esperar por mensagens em canais
		// O select vai esperar até que uma das mensagens chegue
		// e vai executar o código correspondente a essa mensagem
		//Dessa forma com select o programa não espera cada uma das mensagens chegar para seguir o fluxo
		//ele vai enviando as mensagens que vão chegando antes de esperar a próxima
		select {
		case mensagem1 := <-canal1:
			println(mensagem1)
		case mensagem2 := <-canal2:
			println(mensagem2)
		}
		mensagem1 := <-canal1
		println(mensagem1)

		mensagem2 := <-canal2
		println(mensagem2)
	}
}
