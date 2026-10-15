package main

import (
	"fmt"
	"strings"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func MostrarProducto(p Producto) {
	fmt.Printf("Nombre:   %s\n", p.Nombre)
	fmt.Printf("Precio:   $%.2f\n", p.Precio)
	fmt.Printf("Cantidad: %d\n", p.Cantidad)
}

func ValorTotal(p Producto) float64 {
	return p.Precio * float64(p.Cantidad)
}

func TieneStock(p Producto) bool {
	return p.Cantidad > 0
}

func EstadoStock(p Producto) string {
	if TieneStock(p) {
		return "Disponible"
	}
	return "Sin stock"
}

func main() {
	productos := []Producto{
		{Nombre: "Laptop", Precio: 850.00, Cantidad: 5},
		{Nombre: "Mouse inalámbrico", Precio: 15.50, Cantidad: 40},
		{Nombre: "Monitor 24\"", Precio: 179.99, Cantidad: 0},
		{Nombre: "Teclado mecánico", Precio: 65.75, Cantidad: 12},
	}

	separador := strings.Repeat("-", 40)
	valorInventario := 0.0

	fmt.Println("INVENTARIO DE PRODUCTOS")
	fmt.Println(separador)

	for i, p := range productos {
		fmt.Printf("Producto %d\n", i+1)
		MostrarProducto(p)
		total := ValorTotal(p)
		fmt.Printf("Valor total: $%.2f\n", total)
		fmt.Printf("Estado:   %s\n", EstadoStock(p))
		fmt.Println(separador)
		valorInventario += total
	}

	fmt.Printf("Valor total del inventario: $%.2f\n", valorInventario)
}
