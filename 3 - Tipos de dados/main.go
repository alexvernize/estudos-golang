package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Tipos inteiros com sinal
	var i8 int8 = 127                   // -128 a 127
	var i16 int16 = 32767               // -32768 a 32767
	var i32 int32 = 2147483647          // -2147483648 a 2147483647
	var i64 int64 = 9223372036854775807 // -9223372036854775808 a 9223372036854775807
	var i int = 500000                  // depende da arquitetura (32 ou 64 bits)

	// Tipos inteiros sem sinal
	var ui8 uint8 = 255                    // 0 a 255
	var ui16 uint16 = 65535                // 0 a 65535
	var ui32 uint32 = 4294967295           // 0 a 4294967295
	var ui64 uint64 = 18446744073709551615 // 0 a 18446744073709551615
	var ui uint = 500000                   // depende da arquitetura

	// Aliases
	var b byte = 255 // mesmo que uint8
	var r rune = 'A' // mesmo que int32 (representa Unicode)

	// Tipos de ponto flutuante
	var f32 float32 = 3.141592653589793 // precisão de ~6 casas decimais
	var f64 float64 = 3.141592653589793 // precisão de ~15 casas decimais

	// Tipo booleano
	var bo bool = true // true ou false

	// Tipo string
	var s string = "Olá, Go!" // cadeia de caracteres UTF-8

	// Exibindo tipos e tamanhos
	fmt.Println("=== Tipos Inteiros ===")
	fmt.Printf("int8: %d (%d bytes)\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("int16: %d (%d bytes)\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("int32: %d (%d bytes)\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("int64: %d (%d bytes)\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("int: %d (%d bytes)\n", i, unsafe.Sizeof(i))

	fmt.Println("\n=== Tipos Sem Sinal ===")
	fmt.Printf("uint8: %d (%d bytes)\n", ui8, unsafe.Sizeof(ui8))
	fmt.Printf("uint16: %d (%d bytes)\n", ui16, unsafe.Sizeof(ui16))
	fmt.Printf("uint32: %d (%d bytes)\n", ui32, unsafe.Sizeof(ui32))
	fmt.Printf("uint64: %d (%d bytes)\n", ui64, unsafe.Sizeof(ui64))
	fmt.Printf("uint: %d (%d bytes)\n", ui, unsafe.Sizeof(ui))

	fmt.Println("\n=== Aliases ===")
	fmt.Printf("byte: %d (%d bytes)\n", b, unsafe.Sizeof(b))
	fmt.Printf("rune: %U (%d bytes) → '%c'\n", r, unsafe.Sizeof(r), r)

	fmt.Println("\n=== Ponto Flutuante ===")
	fmt.Printf("float32: %.8f (%d bytes)\n", f32, unsafe.Sizeof(f32))
	fmt.Printf("float64: %.15f (%d bytes)\n", f64, unsafe.Sizeof(f64))

	fmt.Println("\n=== Booleano ===")
	fmt.Printf("bool: %t (%d bytes)\n", bo, unsafe.Sizeof(bo))

	fmt.Println("\n=== String ===")
	fmt.Printf("string: %s (%d bytes)\n", s, unsafe.Sizeof(s))
	fmt.Printf("Tamanho da string: %d caracteres\n", len(s))

	// Conversão entre tipos
	fmt.Println("\n=== Conversão de Tipos ===")
	var x int32 = 100
	var y int64 = int64(x) // conversão necessária
	var z float64 = float64(y)

	fmt.Printf("x (int32): %d, y (int64): %d, z (float64): %.2f\n", x, y, z)
}
