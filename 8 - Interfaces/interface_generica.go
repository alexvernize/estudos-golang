/*
Tome cuidado para não utilizar muito e acabar virando uma gambiarra no
seu código
É possível criar uma interface genérica, mas não é recomendado
A interface genérica é uma interface que pode ser utilizada com qualquer tipo
de dado, sem a necessidade de criar uma interface específica para cada tipo
de dado
Isso pode ser útil em alguns casos, mas pode tornar o código mais difícil de entender e manter
Por isso, é importante usar a interface genérica com cautela e apenas quando necessário
*/
package main

import "fmt"

func generica(a interface{}) {
	fmt.Println(a)
}

func main() {
	generica(10)
	generica("Hello")
	generica(true)
	generica(3.14)
	generica([]int{1, 2, 3})
	generica(map[string]int{"a": 1, "b": 2})
	generica(struct {
		nome  string
		idade int
	}{
		nome:  "João",
		idade: 30,
	})
	generica(func() {
		fmt.Println("Função anônima")
	},
	)
	generica(nil)
	generica([]string{"a", "b", "c"})

	//O println é um exemplo de função que aceita qualquer tipo de dado
	//println(10)
	//println("Hello")
	//println(true)
	//println(3.14)
	//println([]int{1, 2, 3})

	//Um exemplo que não é muito bom de uso de interface genérica seria um map
	//map[interface{}]interface{}{
	//	"nome":  "João",
	//	"idade": 30,
	//	"altura": 1.75,
	//	"ativo":  true,
	//	"enderecos": []string{
	//		"Rua A",
	//		"Rua B",
	//	},
}
