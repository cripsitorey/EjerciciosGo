package votos

import "fmt"

func ContarVotos() map[string]int {
	Visitas := make(map[string]int)
	Visitas["Inicio"] = 10
	Visitas["Noticias"] = 20
	Visitas["Deportes"] = 25
	Visitas["Galeria"] = 200
	fmt.Println("Las noticias son: ", Visitas["Noticias"])
	for key, val := range Visitas {
		fmt.Println(key, ":", val)
	}
	plus := 0
	fmt.Println("Escribe la cantidad de Visitas a agregar")
	fmt.Scan(&plus)

	sumador := Visitas["Galeria"] + plus
	Visitas["Galeria"] = sumador
	println("Ahora hay", Visitas["Galeria"], "Visitas a la galeria")
}

func SumarVisitas() {
	visitas := ContarVotos()
	fmt.Println(visitas)

}
