package files

import (
	"bufio"
	"fmt"
	"github/OnlyAdal/godesde0/ejercicios"
	"os"
)

var FileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.GetnumberforMath()
	archivo, err := os.Create(FileName)
	if err != nil {
		fmt.Println("Error al crear el  " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumarTabla() {
	var texto string = ejercicios.GetnumberforMath()
	if !Append(FileName, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(File string, texto string) bool {
	arch, err := os.OpenFile(FileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append ", err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el writte ", err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.ReadFile(FileName)
	if err != nil {
		fmt.Println("Error al leer el archivo ", err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LeoArchivo2() {
	archivo, err := os.Open(FileName)
	if err != nil {
		fmt.Println("Error al leer el archivo ", err.Error())

	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("! " + registro)
	}
	archivo.Close()
}
