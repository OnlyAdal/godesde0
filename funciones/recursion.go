package funciones

func Exponencia(valor int) {
	if valor > 100000000 {
		return

	}
	println(valor)
	Exponencia(valor * 2)
}
