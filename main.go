package main

import (
	"fmt"

	"github.com/Alek151/golangprogra/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
