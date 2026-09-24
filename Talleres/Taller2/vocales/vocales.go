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
		case 'a', 'A':
			voA++
		case 'e', 'E':
			voE++
		case 'i', 'I':
			voI++
		case 'o', 'O':
			voO++
		case 'u', 'U':
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
