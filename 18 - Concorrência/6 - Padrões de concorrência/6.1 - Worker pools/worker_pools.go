/*
O worker pool funciona como se você tivesse uma fila de tarefas a serem executadas
e você teria vários wokers pegando itens dessa fila e executando essas tarefas
de maneira independente e concorrente
*/
package main

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

// Repare que colocando a seta antes do canal, você está dizendo que o canal é somente de leitura
// e colocando depois você está dizendo que o canal é somente de envio
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Criando workers
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	for i := 0; i <= cap(tarefas); i++ {
		tarefas <- i
	}
	close(tarefas) // Fechando o canal de tarefas

	for i := 0; i <= cap(resultados); i++ {
		resultado := <-resultados
		println(resultado)
	}
	close(resultados) // Fechando o canal de resultados
}
