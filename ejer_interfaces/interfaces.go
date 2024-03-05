package ejer_interfaces

import (
	"fmt"
	"github/OnlyAdal/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando ? \n", hu.Sexo())
}
