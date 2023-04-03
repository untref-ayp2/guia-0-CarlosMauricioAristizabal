package main

import "fmt"

func main() {
	numeros := []float32{2.0, 1.0, 5.0, 4.0, 3.0, -3.0, 7.0}
	mostrarPolinomio(numeros)
}

func mostrarPolinomio(numeros []float32) {
	polinomio := ""
	n := 0
	for i := n; i < len(numeros); i++ {
		coeficiente := numeros[i]
		if coeficiente == 0 {
			continue
		}
		signo := "+"
		if coeficiente < 0 {
			signo = "-"
			coeficiente = -coeficiente
		}
		switch i {
		case n:
			polinomio = fmt.Sprintf("%.2f", coeficiente)
		case 1:
			polinomio = fmt.Sprintf("%s %s %.2fX", polinomio, signo, coeficiente)
		case 0:
			polinomio = fmt.Sprintf("%s %s %.2f", polinomio, signo, coeficiente)
		default:
			polinomio = fmt.Sprintf("%s %s %.2fX^%d", polinomio, signo, coeficiente, i)
		}
	}
	fmt.Println(polinomio)
}
