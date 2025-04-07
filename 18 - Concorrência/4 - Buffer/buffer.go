package main

func main() {
	//Criando uma capacidade de buffer para o canal
	canal := make(chan string, 2) // Criando um canal de string com buffer de 2
	//Se eu não usar buffer ele vai ficar esperando até que alguém receba a mensagem
	//e retorna um deadlock porque teria que passar e receber via funções
	canal <- "Olá, mundo!"
	canal <- "Programando em Golang!" // Enviando mensagens para o canal
	//Se eu criar um terceiro canal eu ultrapassaria o buffer e ele não conseguiria enviar
	//a mensagem, o que geraria um deadlock
	//canal <- "Canal com buffer!" // Enviando mensagens para o canal

	mensagem := <-canal
	mensagem2 := <-canal // Recebendo mensagens do canal
	println(mensagem)    // Imprimindo a mensagem recebida do canal
	println(mensagem2)   // Imprimindo a mensagem recebida do canal
}
