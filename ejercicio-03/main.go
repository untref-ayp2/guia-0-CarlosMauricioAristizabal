package main

import "fmt"

func main() {
	lista := []int{5, 2, 6, 4, 3, 10}
	obtenerMenor(lista)
	obtenerMayor(lista)

}

func obtenerMenor(lista []int) {
	menor := lista[0]

	for i := 0; i < len(lista); i++ {
		if menor > lista[i] {
			menor = lista[i]
		}

	}
	fmt.Println("El numero menor es: ", menor)
}
func obtenerMayor(lista []int) {
	mayor := lista[0]

	for i := 0; i < len(lista); i++ {
		if mayor < lista[i] {
			mayor = lista[i]
		}
	}
	fmt.Println("El numero mayor es: ", mayor)
}
