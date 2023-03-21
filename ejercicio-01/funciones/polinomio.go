package funciones

import "fmt"

func mostrarPolinomio(coeficientes []float64) {
	polinomio := ""
	n := len(coeficientes) - 1
	for i := n; i >= 0; i-- {
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
