/*
Interfaces in Go
São tipos que definem um conjunto de métodos.
Um tipo implementa uma interface se implementar todos os métodos dela.
Não é necessário declarar explicitamente que um tipo implementa uma interface.
*/
// Repare que se eu passar um valor do tipo retangulo ou circulo, eu só posso passar um deles para a função
// EscreverArea, pois ela só aceita um tipo
//	func escreverArea(r retangulo) {
//		fmt.Println("Escrevendo área...")
//		fmt.Println(r.largura * r.altura)
//	}
//	func escreverArea(c circulo) {
//		fmt.Println("Escrevendo área...")
//		fmt.Println(3.14 * c.raio * c.raio)
//	}
// Interfaces servem para isso, para que eu possa passar qualquer tipo que implemente a interface
package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	largura float64
	altura  float64
}

type circulo struct {
	raio float64
}

// Criando a interface
type forma interface {
	//Interfaces só tem assinatura de métodos você não pode passar campos da mesma forma da struct
	// Não tem implementação
	area() float64
}

// Criando a função que vai receber a interface
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// Criando os métodos que implementam a interface
func (c circulo) area() float64 {
	//return 3.14 * (c.raio * c.raio)
	return math.Pi * math.Pow(c.raio, 2)
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

func main() {
	// Criando instâncias
	c := circulo{raio: 10}
	r := retangulo{largura: 10, altura: 5}

	// Passando instâncias para a função
	escreverArea(c)
	escreverArea(r)
}
