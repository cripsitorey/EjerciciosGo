package main

import (
	"errors"
	"fmt"
)

type Estudiante struct {
	edad int
}

func (e Estudiante) GetEdad() int {
	return e.edad
}

func (e *Estudiante) SetEdad(edad int) error {
	if edad < 0 || edad < 80 {
		e.edad = edad
		return nil
	} else {
		return errors.New("ERROR, La edad debe estar entre 1 y 79 años")
	}
}

func NewEstudiante(edad int) (*Estudiante, error) {
	estudiante2 := &Estudiante{}
	err := estudiante2.SetEdad(edad)
	if err != nil {
		return nil, err
	}
	return estudiante2, err
}

func main() {
	fmt.Println("Hola bienvenido al sistema de control de edad de estudiantes del curso")
	estudiante1 := Estudiante{}
	age := 0
	fmt.Println("Ingresa la edad del nuevo estudiante")
	fmt.Println("Recuerda que la edad debe estar entre 1 y 79 años")
	fmt.Scan(&age)
	err := estudiante1.SetEdad(age)
	if err != nil {
		fmt.Println("No se puede modificar el registro", err)
	}
	fmt.Println("La edad actual es:", estudiante1.GetEdad())
}
