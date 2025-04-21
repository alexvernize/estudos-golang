package main

import (
	"fmt"
)

type Greeter interface {
	Greet() string
}

type EnglishGreeter struct{}

func (eg EnglishGreeter) Greet() string {
	return "Hello"
}

type SpanishGreeter struct{}

func (eg SpanishGreeter) Greet() string {
	return "!Hola que tal!"
}

func sayHello(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	eg := EnglishGreeter{}
	sg := SpanishGreeter{}

	sayHello(eg)
	sayHello(sg)
}
