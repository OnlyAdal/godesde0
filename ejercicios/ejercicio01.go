package ejercicios

import (
	"strconv"
)

func ConviertoaEntero(numero string) (int, string) {
	var numeroconv int
	numeroconv, err := strconv.Atoi(numero)
	if err != nil {
		return 0, "Hubo un error"
	}

	if numeroconv >= 100 {
		return numeroconv, "Es mayor a 100"

	} else {
		return numeroconv, "Es menor a 100"
	}

}
