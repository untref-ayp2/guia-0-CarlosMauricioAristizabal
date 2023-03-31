package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Println("Ingrese el primer numero entero")
	fmt.Scanln(&a)
	fmt.Println("Ingrese el segundo numero entero")
	fmt.Scanln(&b)

	if b < 0 {
		a = -a
		b = -b
	}
	for i := 0; i < b; i++ {
		c += a
	}
	fmt.Println("El producto es: ", c)
	return
}
