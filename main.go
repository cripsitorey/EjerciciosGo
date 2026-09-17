package main

import "fmt"

func conEs() {
	fmt.Println("Ingrese la cantidad de estudiantes que hay en el curso")
	nEstudiantes := 0
	fmt.Scan(&nEstudiantes)
	notas := 0.0
	for i := 0; i < nEstudiantes; i++ {
		notaE := 0.0
		fmt.Println("Ingresa la nota del estudiante:", i+1)
		fmt.Scan(&notaE)

		notas += notaE
	}
	fmt.Println(averageGrade(nEstudiantes, notas))
	main()

}
func averageGrade(numEst int, grades float64) float64 {
	return grades / float64(numEst)
}

func sum() {
	fmt.Println("Ingresa un numero")
	ene := 0
	fmt.Scan(&ene)
	sumatoria := 0
	for i := 1; i <= ene; i++ {
		sumatoria += i
	}
	fmt.Println(sumatoria)
	main()
}

func CaF() {
	fmt.Println("Ingrese la magnitud que desea convertir a Farenheit")
	celcius := 0.0
	fmt.Scan(&celcius)
	faren := (celcius * 1.8) + 32
	fmt.Println(faren)
	main()
}

func FaC() {
	fmt.Println("Ingrese la magnitud que desea convertir a Celcius")
	faren := 0.0
	fmt.Scan(&faren)
	celcius := (faren - 32) * 5 / 9
	fmt.Println(celcius)
	main()
}

func main() {
	opt := 0
	fmt.Println("0 para salir, 1 para cantidad de estudiantes, 2 para sumatoria, 3 para C a F y 4 para F a C")
	fmt.Scan(&opt)
	switch opt {
	case 0:
		break
	case 1:
		conEs()
	case 2:
		sum()
	case 3:
		CaF()
	case 4:
		FaC()
	default:
		main()
	}

}
