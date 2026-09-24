package vocales

import "fmt"

func Contar() {
	frase := ""
	var voA, voE, voI, voO, voU int
	fmt.Println("Escribe una frase para contar sus vocales: ")
	fmt.Scan(&frase)
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
