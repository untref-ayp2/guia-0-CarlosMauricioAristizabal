package ejercicios

func sumaProductoVector(vector1, vector2 []int) ([]int, int) {
	var sumavector []int
	var producto int

	for i := 0; i < len(vector1); i++ {
		sumavector = append(sumavector, vector1[i]+vector2[i])
		producto += (vector1[i] * vector2[i])
	}
	return sumavector, producto
}
