package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var Numero int
var Err error
var Texto string

func GetnumberforMath() string {

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Ingrese numero: ")
		if scanner.Scan() {
			Numero, Err = strconv.Atoi(scanner.Text())
			if Err == nil {
				for e := 1; e <= 10; e++ {
					multi := Numero * e
					Texto += fmt.Sprintf("%d x %d = %d \n", Numero, e, multi)
				}
				break

			}

		}

	}
	return Texto
}
