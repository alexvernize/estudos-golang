package main

import (
	"fmt"
	"testes/enderecos"
)

// Testa a função de escrita

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua das Américas")
	fmt.Println(tipoEndereco)
}
