package main

import (
	"fmt"
	"runtime"
)

func main() {

	//estado, texto := variables.ConviertoaTexto(1245)
	//fmt.Println(estado)
	//fmt.Println(texto)
	//variables.RestoVariables()
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows")

	} else {
		fmt.Println("Esto es windows")

	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "Darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
