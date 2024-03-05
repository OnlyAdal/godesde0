package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 43, 4, 65, 23, 5, 2, 9, 7}

func MuestroSlice() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]   //slices creado desde un vector desde la posicion
	procion2 := arreglo[:5]  //Slices creado
	porcion3 := arreglo[6:7] //Slice creado desde la posicion 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(procion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("LArgo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)

	}
	fmt.Printf("\nLArgo %d, Capacidad %d", len(nums), cap(nums))

}
