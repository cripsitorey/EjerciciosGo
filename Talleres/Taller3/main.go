package main

import "fmt"

func ejecutarNotas() {
	var notas [6][4]float64

	for i := range notas {
		fmt.Printf("\n--- Estudiante %d ---\n", i+1)
		for j := range notas[i] {
			fmt.Printf("Nota de la materia %d: ", j+1)
			fmt.Scan(&notas[i][j])
		}
	}

	promedios := make([]float64, 0)

	fmt.Println("\nEstudiante | Promedio | Maxima | Minima")

	for i := range notas {
		fila := notas[i][:]

		suma := 0.0
		maxima := fila[0]
		minima := fila[0]

		for _, nota := range fila {
			suma = suma + nota
			if nota > maxima {
				maxima = nota
			}
			if nota < minima {
				minima = nota
			}
		}

		promedio := suma / float64(len(fila))
		promedios = append(promedios, promedio)

		fmt.Printf("%d | %.2f | %.1f | %.1f\n", i+1, promedio, maxima, minima)
	}

	sumaGeneral := 0.0
	for _, p := range promedios {
		sumaGeneral = sumaGeneral + p
	}
	fmt.Printf("Promedio general de la clase: %.2f\n", sumaGeneral/float64(len(promedios)))
}

func main() {
	var opcion int

	fmt.Println("\n===== MENU PRINCIPAL =====")
	fmt.Println("1. Notas de estudiantes")
	fmt.Println("0. Salir")
	fmt.Println("Elige una opcion: ")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		ejecutarNotas()
		main()
	case 0:
		fmt.Println("Hasta luego")
		break
	default:
		fmt.Println("Opcion invalida")
		main()
	}
}
