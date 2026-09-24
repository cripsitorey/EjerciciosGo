package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {
	fmt.Println("Bienvenid@s a la Clase 💕")
	name := ""
	fmt.Println("Escribe tu nombre:")
	fmt.Scan(&name)
	mensaje := saludo.Saludar(name)
	fmt.Println(mensaje)

	operaciones.Menu()
}
