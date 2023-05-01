package main

import (
	"fmt"
	"runtime"

	"github.com/Alek151/golangprogra/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es: ", os)
	case "windows":
		fmt.Println("Esto es: ", os)
	default:
		fmt.Println("Sistema desconocido", os)
	}
}
