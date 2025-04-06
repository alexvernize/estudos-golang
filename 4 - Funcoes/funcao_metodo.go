package main

import (
	"fmt"
	"math"
)

/*
MÉTODOS EM GO

Métodos são funções associadas a tipos específicos.
Eles permitem adicionar comportamentos a tipos personalizados.
*/

// Definindo um tipo personalizado
type Circulo struct {
	raio float64
}

/*
MÉTODO COM RECEIVER POR VALOR

Area é um método do tipo Circulo.
O receiver (c Circulo) aparece antes do nome do método.
Métodos com receiver por valor trabalham com uma cópia do original.
*/
func (c Circulo) Area() float64 {
	return math.Pi * c.raio * c.raio
}

/*
MÉTODO COM RECEIVER POR PONTEIRO

AumentarRaio pode modificar o Circulo original.
Receivers por ponteiro são usados quando:
1. Precisamos modificar o valor original
2. O struct é grande e queremos evitar cópia
*/
func (c *Circulo) AumentarRaio(percentual float64) {
	c.raio *= (1 + percentual/100)
}

/*
MÉTODOS PARA TIPOS BÁSICOS

Podemos definir métodos para tipos personalizados baseados em tipos básicos.
Precisamos primeiro criar um type alias.
*/
type MeuInt int

// Dobro é um método para o tipo MeuInt
func (mi MeuInt) Dobro() MeuInt {
	return mi * 2
}

func main() {
	// USANDO MÉTODOS DE STRUCT
	c := Circulo{raio: 5}

	// Chamando método com receiver por valor
	fmt.Println("Área:", c.Area()) // 78.53981633974483

	// Chamando método com receiver por ponteiro
	c.AumentarRaio(10)                  // Aumenta o raio em 10%
	fmt.Println("Novo raio:", c.raio)   // 5.5
	fmt.Println("Nova área:", c.Area()) // 95.03317777109125

	// MÉTODOS PARA TIPOS BÁSICOS
	num := MeuInt(7)
	fmt.Println("Dobro de", num, "é", num.Dobro()) // 14

	/*
	   OBSERVAÇÕES:
	   1. Go automaticamente converte entre valores e ponteiros ao chamar métodos
	   2. Métodos com receivers por ponteiro podem modificar o valor
	   3. Métodos com receivers por valor não modificam o original
	*/
}
