package main

import (
	"fmt"
	"sort"
)

type Producto struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

type Inventario struct {
	productos map[int]Producto
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) *Producto {
	return &Producto{
		id:       id,
		nombre:   nombre,
		precio:   precio,
		cantidad: cantidad,
	}
}

func (p Producto) GetID() int {
	return p.id
}

func (p Producto) GetNombre() string {
	return p.nombre
}

func (p Producto) GetPrecio() float64 {
	return p.precio
}

func (p Producto) GetCantidad() int {
	return p.cantidad
}

func (p *Producto) SetPrecio(nuevoPrecio float64) {
	p.precio = nuevoPrecio
}

func (p *Producto) SetCantidad(nuevaCantidad int) {
	p.cantidad = nuevaCantidad
}

func (p Producto) MostrarInfo() {
	fmt.Printf("ID: %d | Nombre: %s | Precio: $%.2f | Cantidad: %d\n",
		p.GetID(), p.GetNombre(), p.GetPrecio(), p.GetCantidad())
}

func NuevoInventario() *Inventario {
	return &Inventario{
		productos: make(map[int]Producto),
	}
}

func (inv *Inventario) AgregarProducto(p Producto) {
	if _, existe := inv.productos[p.GetID()]; existe {
		fmt.Printf("Ya existe un producto con ID %d; no se agregó.\n", p.GetID())
		return
	}
	inv.productos[p.GetID()] = p
	fmt.Printf("Producto %q agregado.\n", p.GetNombre())
}

func (inv *Inventario) EliminarProducto(id int) {
	if _, existe := inv.productos[id]; !existe {
		fmt.Printf("No existe un producto con ID %d.\n", id)
		return
	}
	delete(inv.productos, id)
	fmt.Printf("Producto con ID %d eliminado.\n", id)
}

func (inv *Inventario) BuscarProducto(id int) (Producto, bool) {
	p, existe := inv.productos[id]
	return p, existe
}

func (inv *Inventario) ListarProductos() {
	if len(inv.productos) == 0 {
		fmt.Println("El inventario está vacío.")
		return
	}

	ids := make([]int, 0, len(inv.productos))
	for id := range inv.productos {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	fmt.Println("---- Inventario ----")
	for _, id := range ids {
		inv.productos[id].MostrarInfo()
	}
}

func main() {
	inv := NuevoInventario()

	laptop := NuevoProducto(1, "Laptop", 850.00, 5)
	mouse := NuevoProducto(2, "Mouse", 15.50, 30)
	teclado := NuevoProducto(3, "Teclado", 25.00, 20)

	inv.AgregarProducto(*laptop)
	inv.AgregarProducto(*mouse)
	inv.AgregarProducto(*teclado)
	inv.AgregarProducto(*NuevoProducto(2, "Duplicado", 1.00, 1))

	inv.ListarProductos()

	if p, ok := inv.BuscarProducto(2); ok {
		fmt.Print("Encontrado -> ")
		p.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado.")
	}

	if p, ok := inv.BuscarProducto(1); ok {
		p.SetPrecio(799.99)
		p.SetCantidad(8)
		inv.EliminarProducto(p.GetID())
		inv.AgregarProducto(p)
	}

	inv.EliminarProducto(2)
	inv.EliminarProducto(99)

	inv.ListarProductos()
}
