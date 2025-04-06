/*
Métodos são funções associadas a tipos específicos.
Eles permitem adicionar comportamentos a tipos personalizados.
O receiver (c Circulo) aparece antes do nome do método.
Métodos com receiver por valor trabalham com uma cópia do original.
*/

package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func (u usuario) salvar() {
	fmt.Printf("Usuário %s salvo com sucesso!\n", u.nome)
}

func (u usuario) apresentar() {
	fmt.Printf("Olá, meu nome é %s e tenho %d anos.\n", u.nome, u.idade)
}

// Definindo um tipo personalizado
type Circulo struct {
	raio float64
}

// Método com receiver por valor
func (c Circulo) Area() float64 {
	return 3.14 * c.raio * c.raio
}

// Método com receiver por ponteiro
func (c *Circulo) AumentarRaio(percentual float64) {
	c.raio *= (1 + percentual/100)
}

func main() {

	usuario1 := usuario{nome: "João", idade: 25}
	usuario1.fazerAniversario()
	usuario1.fazerAniversario()
	usuario1.salvar()
	usuario1.apresentar()
}
