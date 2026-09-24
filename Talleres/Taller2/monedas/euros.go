package monedas

import "fmt"

func Menu() {
	opt := 0
	dolar := 0.0
	fmt.Println("Ingresa la cantidad en dolares a convertir: ")
	fmt.Scan(&dolar)
	fmt.Println("0 para salir, 1 para convertir a Euro, 2 para Libras Esterlinas, 3 para Wones y 4 para Bitcoin")
	fmt.Scan(&opt)
	switch opt {
	case 0:
		break
	case 1:
		toEuro(dolar)
	case 2:
		toLibra(dolar)
	case 3:
		toWon(dolar)
	case 4:
		toBTC(dolar)
	default:
		Menu()
	}

}

func toEuro(dolares float64) float64 {
	return dolares * 0.88
}

func toBTC(dolares float64) float64 {
	return dolares * 0.000012
}

func toLibra(dolares float64) float64 {
	return dolares * 0.76
}

func toWon(dolares float64) float64 {
	return dolares * 1365.44
}
