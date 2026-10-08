package main

import "fmt"

func descartar() {
	var basura string
	fmt.Scan(&basura)
}

func leerEntero(mensaje string, minimo, maximo int) int {
	for {
		var n int
		fmt.Print(mensaje)
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Entrada inválida. Ingresa un número entero.")
			descartar()
			continue
		}
		if n < minimo || n > maximo {
			fmt.Printf("El valor debe estar entre %d y %d.\n", minimo, maximo)
			continue
		}
		return n
	}
}

func leerNota(mensaje string) float64 {
	for {
		var n float64
		fmt.Print(mensaje)
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("Entrada inválida. Ingresa un número (usa punto decimal).")
			descartar()
			continue
		}
		if n < 0 || n > 10 {
			fmt.Println("La nota debe estar entre 0 y 10.")
			continue
		}
		return n
	}
}

func promedio(valores []float64) float64 {
	suma := 0.0
	for _, v := range valores {
		suma += v
	}
	return suma / float64(len(valores))
}

func maxMin(valores []float64) (alta, baja float64) {
	alta, baja = valores[0], valores[0]
	for _, v := range valores[1:] {
		if v > alta {
			alta = v
		}
		if v < baja {
			baja = v
		}
	}
	return alta, baja
}

func ejecutarNotas() {
	var notas [6][4]float64

	for i := range notas {
		fmt.Printf("\n--- Estudiante %d ---\n", i+1)
		for j := range notas[i] {
			notas[i][j] = leerNota(fmt.Sprintf("Nota de la materia %d (0-10): ", j+1))
		}
	}

	promedios := make([]float64, 0, len(notas))

	fmt.Printf("\n%-12s %-10s %-8s %-8s\n", "Estudiante", "Promedio", "Máxima", "Mínima")
	fmt.Println("------------------------------------------")

	for i := range notas {
		fila := notas[i][:]
		prom := promedio(fila)
		alta, baja := maxMin(fila)
		promedios = append(promedios, prom)

		fmt.Printf("%-12s %-10.2f %-8.1f %-8.1f\n",
			fmt.Sprintf("Est. %d", i+1), prom, alta, baja)
	}

	fmt.Println("------------------------------------------")
	fmt.Printf("Promedio general de la clase: %.2f\n", promedio(promedios))
}

func actividadGanadora(votos map[string]int) (string, int) {
	ganadora := ""
	maxVotos := -1
	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			ganadora = actividad
			maxVotos = cantidad
		}
	}
	return ganadora, maxVotos
}

func ejecutarVotos() {
	opciones := []string{"Deportes", "Videojuegos", "Cine", "Música"}

	votos := map[string]int{}
	for _, actividad := range opciones {
		votos[actividad] = 0
	}

	const totalVotos = 5
	for v := 1; v <= totalVotos; v++ {
		fmt.Printf("\nVoto %d de %d\n", v, totalVotos)
		for i, actividad := range opciones {
			fmt.Printf("  %d. %s\n", i+1, actividad)
		}
		eleccion := leerEntero("Elige una actividad (1-4): ", 1, len(opciones))
		votos[opciones[eleccion-1]]++
	}

	fmt.Println("\n=== RESULTADOS DE LA VOTACIÓN ===")
	for actividad, cantidad := range votos {
		fmt.Printf("%-12s %d voto(s)\n", actividad, cantidad)
	}

	ganadora, cantidad := actividadGanadora(votos)
	fmt.Printf("\nActividad con mayor aceptación: %s (%d voto(s))\n", ganadora, cantidad)
}

func main() {
	var opcion int

	fmt.Println("\n===== MENÚ PRINCIPAL =====")
	fmt.Println("1. Notas de estudiantes")
	fmt.Println("2. Votación de actividades")
	fmt.Println("0. Salir")
	fmt.Print("Elige una opción: ")

	_, err := fmt.Scan(&opcion)
	if err != nil {
		descartar()
		opcion = -1
	}

	switch opcion {
	case 1:
		ejecutarNotas()
		main()
	case 2:
		ejecutarVotos()
		main()
	case 0:
		fmt.Println("Hasta luego.")
		break
	default:
		fmt.Println("Opción inválida. Intenta de nuevo.")
		main()
	}
}
