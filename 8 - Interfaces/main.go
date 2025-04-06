package main

import "fmt"

/*
POLIMORFISMO COM INTERFACES

Interfaces em Go permitem comportamento polimórfico. Um tipo implementa uma interface
automaticamente se implementar todos os seus métodos (duck typing).
*/

// SerVivo define comportamentos comuns para seres vivos
type SerVivo interface {
	Respirar()
	Comunicar() string
}

/*
Pessoa implementa SerVivo implicitamente
(não há declaração explícita de implementação)
*/
type Pessoa struct {
	Nome string
}

// Respirar implementa o método da interface SerVivo
func (p Pessoa) Respirar() {
	fmt.Printf("%s está respirando\n", p.Nome)
}

// Comunicar implementa o método da interface SerVivo
func (p Pessoa) Comunicar() string {
	return "Olá, eu sou " + p.Nome
}

/*
Cachorro também implementa SerVivo
(Diferente struct, mesma interface)
*/
type Cachorro struct {
	Nome string
	Raca string
}

func (c Cachorro) Respirar() {
	fmt.Printf("%s (um %s) está respirando\n", c.Nome, c.Raca)
}

func (c Cachorro) Comunicar() string {
	return "Au Au! Eu sou " + c.Nome
}

// Função que aceita qualquer SerVivo (polimorfismo)
func Apresentar(v SerVivo) {
	v.Respirar()
	fmt.Println(v.Comunicar())
}

func main() {
	// Criando instâncias
	joao := Pessoa{Nome: "João"}
	rex := Cachorro{Nome: "Rex", Raca: "Vira-lata"}

	// Chamada polimórfica
	Apresentar(joao) // Usa a implementação de Pessoa
	Apresentar(rex)  // Usa a implementação de Cachorro

	// Slice de SerVivo (tipos diferentes)
	seres := []SerVivo{joao, rex}
	for _, s := range seres {
		fmt.Println(s.Comunicar()) // Comportamento diferente para cada tipo
	}
}
