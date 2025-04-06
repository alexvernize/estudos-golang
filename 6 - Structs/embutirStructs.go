package main

import "fmt"

/*
MÚLTIPLOS EMBEDDINGS

Um struct pode embuter vários outros structs, combinando seus campos e métodos.
Isso é similar a herança múltipla, mas mais seguro e explícito.
*/

type Motor struct {
	Potencia    int
	Combustivel string
}

func (m Motor) Ligar() {
	fmt.Printf("Motor %s de %dHP ligado!\n", m.Combustivel, m.Potencia)
}

type Bancos struct {
	Material string
	Conforto int
}

func (b Bancos) ConfortoDescricao() string {
	return fmt.Sprintf("Bancos de %s (conforto %d/10)", b.Material, b.Conforto)
}

// Carro combina Motor e Bancos
type Carro struct {
	Motor  // Embedding
	Bancos // Embedding
	Modelo string
}

func main() {
	// Criando um carro com múltiplos embeddings
	carro := Carro{
		Motor: Motor{
			Potencia:    150,
			Combustivel: "Gasolina",
		},
		Bancos: Bancos{
			Material: "Couro",
			Conforto: 8,
		},
		Modelo: "Sedan Premium",
	}

	// Acessando métodos e campos de todos os structs embutidos
	carro.Ligar()                            // Método do Motor
	fmt.Println(carro.ConfortoDescricao())   // Método dos Bancos
	fmt.Printf("Modelo: %s\n", carro.Modelo) // Campo próprio

	// Acessando campos embutidos
	fmt.Printf("Potência: %dHP\n", carro.Potencia) // Promovido de Motor
	fmt.Printf("Material: %s\n", carro.Material)   // Promovido de Bancos
}
