package main

import "fmt"

/*
"SOBRESCRITA" DE MÉTODOS EM GO

Go não tem herança, mas podemos "sobrescrever" métodos de structs embutidos
definindo um método com o mesmo nome no struct container.
*/

// Base - struct pai
type Veiculo struct {
	Marca  string
	Modelo string
}

// Método original
func (v Veiculo) Detalhes() string {
	return fmt.Sprintf("%s %s", v.Marca, v.Modelo)
}

// Carro embute Veiculo
type Carro struct {
	Veiculo
	Portas int
}

// "Sobrescreve" o método Detalhes
func (c Carro) Detalhes() string {
	// Ainda podemos acessar o método original se necessário
	base := c.Veiculo.Detalhes()
	return fmt.Sprintf("%s com %d portas", base, c.Portas)
}

func main() {
	// Criando um carro
	gol := Carro{
		Veiculo: Veiculo{
			Marca:  "VW",
			Modelo: "Gol",
		},
		Portas: 4,
	}

	// Chamando o método "sobrescrito"
	fmt.Println(gol.Detalhes()) // "VW Gol com 4 portas"

	// Ainda podemos acessar o método original explicitamente
	fmt.Println(gol.Veiculo.Detalhes()) // "VW Gol"
}
