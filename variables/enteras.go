package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Int comun = ", intcomun)
	fmt.Println("Int comun de 32 = ", intde32)
	fmt.Println("Int comun de 64 = ", intde64)

}
