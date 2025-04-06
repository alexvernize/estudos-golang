package main

import "fmt"

func main() {
	/*
	   OPERADORES ARITMÉTICOS:
	   Realizam operações matemáticas básicas
	*/
	a := 10
	b := 3

	fmt.Println("\n=== OPERADORES ARITMÉTICOS ===")
	fmt.Println("a + b =", a+b) // Soma: 13
	fmt.Println("a - b =", a-b) // Subtração: 7
	fmt.Println("a * b =", a*b) // Multiplicação: 30
	fmt.Println("a / b =", a/b) // Divisão inteira: 3
	fmt.Println("a % b =", a%b) // Resto da divisão: 1

	// Divisão com números float
	fmt.Println("a / b (float) =", float64(a)/float64(b)) // 3.333...

	/*
	   OPERADORES DE ATRIBUIÇÃO:
	   Atribuem valores a variáveis
	*/
	fmt.Println("\n=== OPERADORES DE ATRIBUIÇÃO ===")
	x := 5
	x += 3                     // Equivalente a x = x + 3
	fmt.Println("x += 3 →", x) // 8

	x -= 2                     // x = x - 2
	fmt.Println("x -= 2 →", x) // 6

	x *= 4                     // x = x * 4
	fmt.Println("x *= 4 →", x) // 24

	x /= 3                     // x = x / 3
	fmt.Println("x /= 3 →", x) // 8

	x %= 5                     // x = x % 5
	fmt.Println("x %= 5 →", x) // 3

	/*
	   OPERADORES DE COMPARAÇÃO:
	   Comparam valores e retornam um booleano (true/false)
	*/
	fmt.Println("\n=== OPERADORES DE COMPARAÇÃO ===")
	fmt.Println("a == b →", a == b) // Igualdade: false
	fmt.Println("a != b →", a != b) // Diferença: true
	fmt.Println("a < b →", a < b)   // Menor que: false
	fmt.Println("a > b →", a > b)   // Maior que: true
	fmt.Println("a <= b →", a <= b) // Menor ou igual: false
	fmt.Println("a >= b →", a >= b) // Maior ou igual: true

	/*
	   OPERADORES LÓGICOS:
	   Trabalham com valores booleanos
	*/
	fmt.Println("\n=== OPERADORES LÓGICOS ===")
	verdadeiro := true
	falso := false

	fmt.Println("verdadeiro && falso →", verdadeiro && falso) // AND: false
	fmt.Println("verdadeiro || falso →", verdadeiro || falso) // OR: true
	fmt.Println("!verdadeiro →", !verdadeiro)                 // NOT: false

	/*
	   OPERADORES DE BIT:
	   Realizam operações bit a bit
	*/
	fmt.Println("\n=== OPERADORES DE BIT ===")
	num1 := 10 // 1010 em binário
	num2 := 3  // 0011 em binário

	fmt.Println("num1 & num2 →", num1&num2)   // AND bit a bit: 2 (0010)
	fmt.Println("num1 | num2 →", num1|num2)   // OR bit a bit: 11 (1011)
	fmt.Println("num1 ^ num2 →", num1^num2)   // XOR bit a bit: 9 (1001)
	fmt.Println("num1 &^ num2 →", num1&^num2) // AND NOT: 8 (1000)
	fmt.Println("num1 << 2 →", num1<<2)       // Deslocamento à esquerda: 40 (101000)
	fmt.Println("num1 >> 1 →", num1>>1)       // Deslocamento à direita: 5 (0101)

	/*
	   OPERADORES DE MEMÓRIA:
	   Trabalham com ponteiros
	*/
	fmt.Println("\n=== OPERADORES DE MEMÓRIA ===")
	valor := 42
	ptr := &valor // & obtém o endereço de memória

	fmt.Println("Valor:", valor)             // 42
	fmt.Println("Endereço:", ptr)            // 0xc0000180a8 (exemplo)
	fmt.Println("Valor via ponteiro:", *ptr) // * acessa o valor no endereço

	*ptr = 100                        // Modificando o valor através do ponteiro
	fmt.Println("Novo valor:", valor) // 100

	/*
	   OPERADORES ESPECIAIS EM GO:
	   Operadores úteis para tipos específicos
	*/
	fmt.Println("\n=== OPERADORES ESPECIAIS ===")

	// Operador de concatenação de strings
	str1 := "Hello"
	str2 := "World"
	fmt.Println(str1 + " " + str2) // Hello World

	// Operador de channel (veremos em concorrência)
	ch := make(chan int)
	go func() { ch <- 42 }()
	fmt.Println("Valor do channel:", <-ch) // 42

	// Operador de tipo (type assertion)
	var interfaceVar interface{} = "Go é incrível"
	str, ok := interfaceVar.(string)
	fmt.Println("Type assertion:", str, ok) // Go é incrível true

	/*
	   PRECEDÊNCIA DE OPERADORES:
	   Ordem de avaliação quando vários operadores aparecem juntos
	*/
	fmt.Println("\n=== PRECEDÊNCIA ===")
	resultado := 5 + 3*2 - 1/2
	fmt.Println("5 + 3*2 - 1/2 =", resultado) // 10 (3*2=6, 1/2=0, 5+6-0=11)

	// Use parênteses para alterar a precedência
	resultado = (5 + 3) * (2 - 1) / 2
	fmt.Println("(5 + 3)*(2 - 1)/2 =", resultado) // 4
}
