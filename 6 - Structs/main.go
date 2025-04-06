package main

import (
	"fmt"
	"time"
)

/*
STRUCTS EM GO

Structs são tipos de dados que agrupam campos relacionados.
Eles são similares a classes em outras linguagens (mas sem herança).
*/

// 1. DEFINIÇÃO BÁSICA DE STRUCT
type Pessoa struct {
	Nome       string
	Idade      int
	Altura     float64
	Ativo      bool
	Nascimento time.Time
}

// 2. STRUCT COM TAGS (para JSON, validação, etc.)
type Usuario struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"min=3"`
	Senha    string `json:"-"` // O - omite o campo na serialização
}

// 3. STRUCTS ANINHADOS
type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Pessoa              // Campos embutidos (embedded)
	Endereco   Endereco // Campo nomeado
	Telefone   string
	Assinatura bool
}

// 4. MÉTODOS PARA STRUCTS
func (p Pessoa) Saudacao() string {
	return fmt.Sprintf("Olá, meu nome é %s e tenho %d anos!", p.Nome, p.Idade)
}

// Método com receiver de ponteiro para modificar o struct
func (p *Pessoa) Aniversario() {
	p.Idade++
}

// 5. CONSTRUTORES (convenção, não é built-in)
func NovaPessoa(nome string, idade int, altura float64) Pessoa {
	return Pessoa{
		Nome:   nome,
		Idade:  idade,
		Altura: altura,
		Ativo:  true,
	}
}

func main() {
	// ==========================================
	// INSTANCIANDO STRUCTS
	// ==========================================

	// Forma 1: Declaração com zero values
	var p1 Pessoa
	fmt.Println("Pessoa 1 (zero values):", p1)

	// Forma 2: Literal com todos os campos
	p2 := Pessoa{
		Nome:       "João",
		Idade:      30,
		Altura:     1.75,
		Ativo:      true,
		Nascimento: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	fmt.Println("Pessoa 2 (literal completo):", p2)

	// Forma 3: Literal parcial (campos omitidos terão zero values)
	p3 := Pessoa{Nome: "Maria", Idade: 25}
	fmt.Println("Pessoa 3 (literal parcial):", p3)

	// Forma 4: Usando new (retorna ponteiro)
	p4 := new(Pessoa)
	p4.Nome = "Carlos"
	p4.Idade = 40
	fmt.Println("Pessoa 4 (com new):", *p4)

	// Forma 5: Usando construtor
	p5 := NovaPessoa("Ana", 28, 1.68)
	fmt.Println("Pessoa 5 (construtor):", p5)

	// ==========================================
	// ACESSANDO CAMPOS
	// ==========================================
	fmt.Println("\nNome da Pessoa 2:", p2.Nome)
	p2.Idade = 31 // Modificando um campo
	fmt.Println("Pessoa 2 atualizada:", p2)

	// ==========================================
	// STRUCTS ANINHADOS
	// ==========================================
	cliente := Cliente{
		Pessoa: Pessoa{
			Nome:   "Fernanda",
			Idade:  35,
			Altura: 1.62,
		},
		Endereco: Endereco{
			Rua:    "Av. Principal",
			Numero: 123,
			Cidade: "São Paulo",
			Estado: "SP",
		},
		Telefone:   "11-99999-8888",
		Assinatura: true,
	}

	// Acessando campos aninhados
	fmt.Println("\nCliente:")
	fmt.Println("Nome:", cliente.Nome)        // Campo embutido (embedded)
	fmt.Println("Rua:", cliente.Endereco.Rua) // Campo nomeado
	fmt.Println("Telefone:", cliente.Telefone)

	// ==========================================
	// MÉTODOS
	// ==========================================
	fmt.Println("\nMétodos:")
	fmt.Println(p2.Saudacao())
	p2.Aniversario()
	fmt.Println("Depois do aniversário:", p2.Saudacao())

	// ==========================================
	// COMPARANDO STRUCTS
	// ==========================================
	p6 := Pessoa{Nome: "João", Idade: 30}
	p7 := Pessoa{Nome: "João", Idade: 30}
	fmt.Println("\nComparando structs:")
	fmt.Println("p6 == p7:", p6 == p7) // true (campos com mesmos valores)

	// ==========================================
	// STRUCTS E PONTEIROS
	// ==========================================
	ptr := &p3
	fmt.Println("\nAcessando via ponteiro:")
	fmt.Println("Nome via ptr:", ptr.Nome) // Go faz dereferência automaticamente
	ptr.Idade = 26
	fmt.Println("Idade modificada via ptr:", p3.Idade)

	// ==========================================
	// STRUCTS ANÔNIMOS
	// ==========================================
	anon := struct {
		Tipo  string
		Valor int
	}{
		Tipo:  "Temporário",
		Valor: 100,
	}
	fmt.Println("\nStruct anônimo:", anon)
}
