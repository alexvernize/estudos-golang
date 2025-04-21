// TESTE DE UNIDADE
// Testa uma pequena parte do código
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	/*
		// Testa se o tipo de endereço é "rua"
		enderecoParaTeste := "Rua das Flores"
		tipoDeEnderecoEsperado := "rua"
		tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
		if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
			t.Errorf("Esperado '%s', mas obteve '%s'", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
		}
	*/
	// Testa se o tipo de endereço é "rua"
	tipo := TipoDeEndereco("Rua das Flores")
	if tipo != "rua" {
		t.Errorf("Esperado 'rua', mas obteve '%s'", tipo)
	}

	// Testa se o tipo de endereço é "avenida"
	tipo = TipoDeEndereco("Avenida Brasil")
	if tipo != "avenida" {
		t.Errorf("Esperado 'avenida', mas obteve '%s'", tipo)
	}

	// Testa se o tipo de endereço é inválido
	tipo = TipoDeEndereco("Casa Verde")
	if tipo != "Tipo inválido" {
		t.Errorf("Esperado 'Tipo inválido', mas obteve '%s'", tipo)
	}
}
