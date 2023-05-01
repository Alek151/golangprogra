package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CrearTabla() {
	scanner := bufio.NewScanner(os.Stdin)

	var numero int
	var err error

	for {
		fmt.Println("Ingrese el número")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 0; i <= 10; i++ {
		println(numero, "*", i, "=", numero*i)
	}
}
