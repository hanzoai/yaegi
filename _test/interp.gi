package main

import (
	"github.com/hanzoai/yaegi/interp"
)

func main() {
	i := interp.New(interp.Opt{})
	i.Eval(`println("Hello")`)
}

// Output:
// Hello
