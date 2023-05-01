package files

import (
	//"bufio"
	"fmt"
	//"ioutil"
	"os"

	"github.com/Alek151/golangprogra/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.CrearTabla()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo: " + err.Error())
		return
	}

	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.CrearTabla()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append: " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
	}

	arch.Close()

	return true

}

func LeoArchivo() {}
