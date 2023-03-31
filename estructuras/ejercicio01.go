package main

import "fmt"

func main() {
	var numero int
	fmt.Println("ingrese un numero")
	fmt.Scanln(&numero)

	if numero < 1 {
		fmt.Println("El numero debe ser positivo")
		return
	} else {
		resultado := 1

		for i := numero; i > 0; i-- {
			resultado = resultado * i
		}
		fmt.Println("El factorial es", resultado)
	}
}
