package ejercicios

func sumaDeArreglo(arreglo []int) int {
	var suma int
	for i := 0; i < len(arreglo); i++ {
		suma += arreglo[i]
	}
	return suma
}
