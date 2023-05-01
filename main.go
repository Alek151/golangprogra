package main

import (
	"fmt"

	"github.com/Alek151/golangprogra/ejercicios"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es: ", os)
	case "windows":
		fmt.Println("Esto es: ", os)
	default:
		fmt.Println("Sistema desconocido", os)
	}*/
	numero, texto := ejercicios.DevuelveVariables("500")
	fmt.Println(numero)
	fmt.Println(texto)
}
