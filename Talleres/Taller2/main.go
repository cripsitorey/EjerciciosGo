package main

import (
	"converter/monedas"
	"converter/vocales"
	"fmt"
)

func main() {
	opt := 0
	fmt.Println("1 para Conversor de monedas, 2 para contador de vocales y 0 para salir")
	fmt.Scan(&opt)
	switch opt {
	case 0:
		break
	case 1:
		monedas.Menu()
		main()
	case 2:
		vocales.Contar()
		main()
	default:
		main()
	}
}
