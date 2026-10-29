package main

import "fmt"

type Persona struct {
	nombre string
}

func (p Persona) GetNombre() string {
	return p.nombre
}

func (p *Persona) SetNombre(nombre string) {
	p.nombre = nombre
}

func main() {
	name := ""
	persona := Persona{}
	fmt.Println("Ingresa su nombre")
	fmt.Scan(&name)
	persona.SetNombre(name)
	fmt.Println("Tu nombre es:")
	fmt.Println(persona.GetNombre())
	opt := 0
	fmt.Println("Si quieres cambiar tu nombre selecciona 1, si no pon 0 para salir")
	fmt.Scan(&opt)
	if opt == 1 {
		fmt.Scan(&name)
		persona.SetNombre(name)
		fmt.Println("Tu nuevo nombre es:", persona.GetNombre())
	} else {
		fmt.Println("Chao")
	}
}
