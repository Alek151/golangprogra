package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CrearTabla() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese el número")
	if scanner.Scan() {

		numero, err := strconv.Atoi(scanner.Text())

		if err != nil {
			panic("Ingresaste un dato incorrecto: " + err.Error())
		}

		for i := 0; i <= 10; i++ {
			println(numero, "*", i, "=", numero*i)
		}
	}

}
