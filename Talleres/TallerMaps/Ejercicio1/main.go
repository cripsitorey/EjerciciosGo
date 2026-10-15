package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func separarPalabras(texto string) []string {
	texto = strings.ToLower(texto)
	return strings.FieldsFunc(texto, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func contarPalabras(palabras []string) map[string]int {
	frecuencias := make(map[string]int)
	for _, palabra := range palabras {
		frecuencias[palabra]++
	}
	return frecuencias
}

func palabrasMasRepetidas(frecuencias map[string]int) (int, []string) {
	maximo := 0
	var ganadoras []string
	for palabra, cantidad := range frecuencias {
		if cantidad > maximo {
			maximo = cantidad
			ganadoras = []string{palabra}
		} else if cantidad == maximo {
			ganadoras = append(ganadoras, palabra)
		}
	}
	sort.Strings(ganadoras)
	return maximo, ganadoras
}

func main() {
	lector := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una frase o texto: ")
	texto, err := lector.ReadString('\n')
	if err != nil && texto == "" {
		fmt.Println("Error al leer la entrada:", err)
		return
	}

	palabras := separarPalabras(texto)
	if len(palabras) == 0 {
		fmt.Println("No se ingresaron palabras válidas.")
		return
	}

	frecuencias := contarPalabras(palabras)

	claves := make([]string, 0, len(frecuencias))
	for palabra := range frecuencias {
		claves = append(claves, palabra)
	}
	sort.Strings(claves)

	fmt.Println("\nFrecuencia de palabras:")
	fmt.Println("-----------------------")
	for _, palabra := range claves {
		fmt.Printf("%-20s %d\n", palabra, frecuencias[palabra])
	}
	fmt.Printf("\nTotal de palabras: %d | Palabras distintas: %d\n",
		len(palabras), len(frecuencias))

	maximo, ganadoras := palabrasMasRepetidas(frecuencias)
	if len(ganadoras) == 1 {
		fmt.Printf("\nLa palabra que más se repite es \"%s\" (%d veces).\n",
			ganadoras[0], maximo)
	} else {
		fmt.Printf("\nHay un empate entre las palabras más repetidas (%d veces cada una): %s\n",
			maximo, strings.Join(ganadoras, ", "))
	}
}
