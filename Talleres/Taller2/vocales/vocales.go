package vocales

import (
	"bufio"
	"fmt"
	"os"
)

func Contar() {
	var voA, voE, voI, voO, voU int
	fmt.Println("Escribe una frase para contar sus vocales: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()
	for _, v := range frase {
		switch v {
		case 'a', 'A', 'á', 'Á':
			voA++
		case 'e', 'E', 'é', 'É':
			voE++
		case 'i', 'I', 'í', 'Í':
			voI++
		case 'o', 'O', 'ó', 'Ó':
			voO++
		case 'u', 'U', 'ú', 'Ú', 'ü', 'Ü':
			voU++
		}
	}
	fmt.Println("La frase que pusiste tiene un total de: ")
	fmt.Println(voA, " As")
	fmt.Println(voE, " Es")
	fmt.Println(voI, " Is")
	fmt.Println(voO, " Os")
	fmt.Println(voU, " Us")
}
