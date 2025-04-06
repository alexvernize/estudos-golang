package matematica // Define o nome do pacote

// Soma calcula a soma de dois números inteiros
// Por começar com letra maiúscula, esta função é exportada (pública)
func Soma(a, b int) int {
	return a + b
}

// subtracao calcula a subtração entre dois números
// Por começar com letra minúscula, esta função NÃO é exportada (privada)
func subtracao(a, b int) int {
	return a - b
}

// Constante pública (exportada)
const Pi = 3.14159265359

// variável privada (não exportada)
var versao = "1.0"
