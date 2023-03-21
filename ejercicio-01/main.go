package main

import "fmt"

func main() {
	coeficientes := []float64{2.0, 1.0, 5.0, 4.0, 3.0, -3.0, 7.0}
	mostrarPolinomio(coeficientes)
}

func mostrarPolinomio(coeficientes []float64) {
	polinomio := ""
	n := 0
	for i := n; i < len(coeficientes); i++ {
		coef := coeficientes[i]
		if coef == 0 {
			continue
		}
		signo := "+"
		if coef < 0 {
			signo = "-"
			coef = -coef
		}
		switch i {
		case n:
			polinomio = fmt.Sprintf("%.2f", coef)
		case 1:
			polinomio = fmt.Sprintf("%s %s %.2fX", polinomio, signo, coef)
		case 0:
			polinomio = fmt.Sprintf("%s %s %.2f", polinomio, signo, coef)
		default:
			polinomio = fmt.Sprintf("%s %s %.2fX^%d", polinomio, signo, coef, i)
		}
	}
	fmt.Println(polinomio)
}
