package main

import "fmt"

/*
COMPOSIÇÃO EM GO (EMBEDDING)

Em Go não temos herança tradicional, mas sim composição através de structs embutidos.
Isso significa que um struct pode "herdar" campos e métodos de outro struct
por embutimento (embedding), sem criar uma relação de pai/filho.
*/

// Animal define campos e métodos básicos para um animal
type Animal struct {
	Nome  string
	Idade int
}

// Comer é um método disponível para qualquer Animal
func (a Animal) Comer() {
	fmt.Printf("%s está comendo...\n", a.Nome)
}

/*
Cachorro embute Animal (composição)
Isso significa que:
- Todos os campos de Animal são promovidos para Cachorro
- Todos os métodos de Animal estão disponíveis em Cachorro
*/
type Cachorro struct {
	Animal // Embedding (sem nome de campo = promoção automática)
	Raca   string
}

// Latir é um método específico de Cachorro
func (c Cachorro) Latir() {
	fmt.Printf("%s diz: Au Au!\n", c.Nome) // c.Nome vem do Animal embutido
}

func main() {
	// Criando um Cachorro
	rex := Cachorro{
		Animal: Animal{ // Inicializando o struct embutido
			Nome:  "Rex",
			Idade: 3,
		},
		Raca: "Vira-lata",
	}

	// Acessando campos e métodos promovidos
	fmt.Printf("%s é um %s de %d anos\n",
		rex.Nome,  // Campo promovido de Animal
		rex.Raca,  // Campo direto de Cachorro
		rex.Idade) // Campo promovido de Animal

	rex.Comer() // Método promovido de Animal
	rex.Latir() // Método específico de Cachorro
}
