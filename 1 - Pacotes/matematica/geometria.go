package matematica // Mesmo pacote que aritmetica.go

// AreaRetangulo calcula a área de um retângulo
func AreaRetangulo(base, altura float64) float64 {
	return base * altura
}

// AreaCirculo calcula a área de um círculo
// Note que podemos usar a constante Pi definida em aritmetica.go
// pois estão no mesmo pacote
func AreaCirculo(raio float64) float64 {
	return Pi * raio * raio
}
