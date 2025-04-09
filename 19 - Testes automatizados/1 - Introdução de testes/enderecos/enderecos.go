package enderecos

import (
	"strings"
)

// TipoDeEndereco recebe um endereço, verifica se ele é válido e retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "praça", "alameda", "travessa", "largo", "vila",
		"bairro", "conjunto", "setor", "quadra", "loteamento"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)                        // Converte para minúsculo
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] // Pega a primeira palavra do endereço por isso o índice 0

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
			return tipo
		}
	}
	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}
	return "Tipo inválido"
}
