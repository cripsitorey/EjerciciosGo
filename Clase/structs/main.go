package main

import (
	"bufio"
	"fmt"
	"os"
)

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	if nuevoPrecio < 0 {
		fmt.Println("El precio no puede ser negativo")
		return
	}
	p.Precio = nuevoPrecio
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	if nuevaCantidad < 0 {
		fmt.Println("La cantidad no puede ser negativa")
		return
	}
	p.Cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d\n", p.ID, p.Nombre, p.Precio, p.Cantidad)
}

func (inv *Inventario) AgregarProducto(p Producto) {
	inv.productos[p.ID] = p
}

func (inv *Inventario) EliminarProducto(id int) {
	_, existe := inv.productos[id]
	if !existe {
		fmt.Println("No existe un producto con ese ID")
		return
	}
	delete(inv.productos, id)
	fmt.Println("Producto eliminado")
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv *Inventario) ListarProductos() {
	if len(inv.productos) == 0 {
		fmt.Println("El inventario esta vacio")
		return
	}
	for _, p := range inv.productos {
		p.MostrarInfo()
	}
}

func calcularValorInventario(p ...Producto) float64 {
	total := 0.0
	for _, p := range p {
		total += p.Precio * float64(p.Cantidad)
	}
	return total
}

func leerLinea(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func leerEntero(scanner *bufio.Scanner) int {
	n := 0
	fmt.Sscan(leerLinea(scanner), &n)
	return n
}

func leerDecimal(scanner *bufio.Scanner) float64 {
	n := 0.0
	fmt.Sscan(leerLinea(scanner), &n)
	return n
}

func main() {
	inventario := Inventario{productos: map[int]Producto{}}
	scanner := bufio.NewScanner(os.Stdin)

	product1 := NuevoProducto(1, "Nexa 35k puffs", 7, 3)

	product2 := Producto{
		ID:       2,
		Nombre:   "Spaceman 50k puffs",
		Precio:   16,
		Cantidad: 1,
	}

	product3 := Producto{
		3,
		"Vint 50k puffs",
		23,
		0,
	}

	inventario.AgregarProducto(product1)
	inventario.AgregarProducto(product2)
	inventario.AgregarProducto(product3)

	opt := -1
	for opt != 0 {
		fmt.Println()
		fmt.Println("Escribe el numero de la opcion:")
		fmt.Println("1: Listar productos")
		fmt.Println("2: Buscar producto")
		fmt.Println("3: Agregar producto")
		fmt.Println("4: Eliminar producto")
		fmt.Println("5: Actualizar precio")
		fmt.Println("6: Actualizar cantidad")
		fmt.Println("7: Ver valor total del inventario")
		fmt.Println("0: Salir")
		opt = leerEntero(scanner)

		switch opt {
		case 1:
			inventario.ListarProductos()
		case 2:
			fmt.Println("ID del producto:")
			id := leerEntero(scanner)
			p, existe := inventario.BuscarProducto(id)
			if existe {
				p.MostrarInfo()
				fmt.Println("Valor en inventario:", calcularValorInventario(p))
			} else {
				fmt.Println("Producto no encontrado")
			}
		case 3:
			fmt.Println("ID del producto:")
			id := leerEntero(scanner)
			_, existe := inventario.BuscarProducto(id)
			if existe {
				fmt.Println("Ya existe un producto con ese ID")
				break
			}
			fmt.Println("Nombre:")
			nombre := leerLinea(scanner)
			fmt.Println("Precio:")
			precio := leerDecimal(scanner)
			fmt.Println("Cantidad:")
			cantidad := leerEntero(scanner)
			if precio < 0 || cantidad < 0 {
				fmt.Println("Precio y cantidad no pueden ser negativos")
				break
			}
			inventario.AgregarProducto(NuevoProducto(id, nombre, precio, cantidad))
			fmt.Println("Producto agregado")
		case 4:
			fmt.Println("ID del producto a eliminar:")
			id := leerEntero(scanner)
			inventario.EliminarProducto(id)
		case 5:
			fmt.Println("ID del producto:")
			id := leerEntero(scanner)
			p, existe := inventario.BuscarProducto(id)
			if !existe {
				fmt.Println("Producto no encontrado")
				break
			}
			fmt.Println("Nuevo precio:")
			p.ActualizarPrecio(leerDecimal(scanner))
			inventario.AgregarProducto(p)
			p.MostrarInfo()
		case 6:
			fmt.Println("ID del producto:")
			id := leerEntero(scanner)
			p, existe := inventario.BuscarProducto(id)
			if !existe {
				fmt.Println("Producto no encontrado")
				break
			}
			fmt.Println("Nueva cantidad:")
			p.ActualizarCantidad(leerEntero(scanner))
			inventario.AgregarProducto(p)
			p.MostrarInfo()
		case 7:
			productos := []Producto{}
			for _, p := range inventario.productos {
				productos = append(productos, p)
			}
			fmt.Println("Valor total del inventario:", calcularValorInventario(productos...))
		case 0:
			fmt.Println("Has salido")
		default:
			fmt.Println("Opcion no valida")
		}
	}
}
