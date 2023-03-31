package main

import "fmt"

func main() {
	arreglo := []int{1, 5, 5, 4, 3, 2}
	fmt.Println(sumaDeArreglo(arreglo))
}

func sumaDeArreglo(arreglo []int) int {
	var suma int
	for i := 0; i < len(arreglo); i++ {
		suma += arreglo[i]
	}
	return suma
}
