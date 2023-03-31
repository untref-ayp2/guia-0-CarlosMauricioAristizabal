package main

import "fmt"

func main() {
	var a int
	fmt.Println("Ingrese un numero entero")
	fmt.Scanln(&a)

	primo := true
	for i := 3; i < a && primo; i++ {
		if a%i != 0 {
			primo = false
		}
	}
	fmt.Println("El numero es primo: ", primo)
}
