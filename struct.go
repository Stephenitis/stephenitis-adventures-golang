package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs")
	p := pair{"Hello", "World"}
	p.Say()
}

type pair struct {
	variable1, variable2 string
}

func (p *pair) Say() {
	fmt.Println(p.variable1 + p.variable2)
}
