package main

import "fmt"

func main() {

	//println(opciones(5))
	opciones2()
}

func opciones(numero int) (respuesta string) {

	switch numero {
	case 1:
		respuesta = "Ha elegido la opcion 1"
	case 2:
		respuesta = "Ha elegido la opcion 2"
	case 3:
		respuesta = "Ha elegido la opcion 3"
	case 4:
		respuesta = "Ha elegido la opcion 4"
	default:
		respuesta = "Ha seleccionado una opcion invalida"
	}
	return respuesta

}

func opciones2() {

	println("Seleccione una opcion")
	println("Opcion 1")
	println("Opcion 2")
	println("Opcion 3")
	println("Opcion 4")

	var eleccion int
	fmt.Scanln(&eleccion)

	if eleccion > 0 && eleccion < 5 {
		fmt.Print("Su eleccion es: ", eleccion)
	} else {
		fmt.Println("Ha seleccionado una opcion invalida")
	}
}
