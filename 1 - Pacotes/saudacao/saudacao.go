package saudacao

import "fmt"

// Ola imprime uma saudação genérica
func Ola() {
	fmt.Println("Olá, seja bem-vindo!")
}

// OlaNome imprime uma saudação personalizada
func OlaNome(nome string) {
	fmt.Printf("Olá, %s! Seja bem-vindo!\n", nome)
}

// init é uma função especial que é executada quando o pacote é importado
func init() {
	fmt.Println("Pacote de saudação inicializado!")
}
