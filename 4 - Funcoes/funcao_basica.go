package main

import "fmt"

// Função simples sem parâmetros
func saudacao() {
	fmt.Println("Bem-vindo ao estudo de funções em Go!")
}

// Função com parâmetro
func saudar(nome string) {
	fmt.Printf("Olá, %s!\n", nome)
}

// Função com retorno
func soma(a, b int) int {
	return a + b
}

// Função com múltiplos retornos
func calculos(a, b int) (int, int) {
	soma := a + b
	diferenca := a - b
	return soma, diferenca
}

func main() {
	saudacao()
	saudar("João")

	resultado := soma(5, 3)
	fmt.Println("5 + 3 =", resultado)

	s, d := calculos(10, 4)
	fmt.Printf("Soma: %d, Diferença: %d\n", s, d)
}
