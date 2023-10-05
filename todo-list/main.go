package main

import "fmt"

//Estructura
type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi Nombre es ", p.nombre)
}

func main() {
	var a [5]int
	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50
	fmt.Println(a)
	var a1 = [5]int{11, 12, 13, 14, 15}
	fmt.Println(a1)
	var a2 = [...]int{11, 12, 13, 14, 15, 16, 17}
	fmt.Println(a2[3])

	//Iterar una matriz con For
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}

	//iterar una matriz con Range
	for index, value := range a {
		fmt.Printf("Indice: %d, Valor : %d \n", index, value)
	}

	// Matriz Multidimensional
	var matriz1 = [3][3]int{{1, 11, 12}, {10, 1, 23}, {534, 11, 1}}
	fmt.Println(matriz1)

	/// Rebanadas
	diasSmana := []string{"Doming", "Lunes", "Martes", "Miecoles", "Jueves", "viernes", "Sabado"}
	fmt.Println(diasSmana)

	diaRebanada := diasSmana[0:5]
	fmt.Println(diaRebanada)
	diaRebanada = append(diaRebanada, "Viernes", "Sabado", "Otro dia ")
	fmt.Println(diaRebanada)
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)
	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	// Funcion Make Crea rebanadas
	nombres := make([]string, 5, 10)
	nombres[0] = "Jonh"
	fmt.Println(nombres)

	// Funcion Copy Copiar rebanadas
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
	copy(rebanada2, rebanada1)
	fmt.Println("Cantidad Elemntos", copy(rebanada2, rebanada1))
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)

	//Mapas
	colors := map[string]string{
		"rojo":  "#FF0000",
		"azul":  "#00FF00",
		"verde": "#0000FF",
	}
	fmt.Println(colors)
	fmt.Println(colors["rojo"])
	colors["negro"] = "#000000"
	fmt.Println(colors)
	if valor, ok := colors["verde"]; ok {
		fmt.Println("Si existe la clave")
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave")

	}

	//Eliminar elemento mapa
	delete(colors, "azul")
	fmt.Println(colors)
	//Iterar Mapa
	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}

	//Estructura de Datos
	var p Persona
	p.nombre = "Jonh"
	p.edad = 30
	p.correo = "jn@correo.com"
	fmt.Println(p)
	p1 := Persona{"Jonh Niño", 30, "nino@correo.com"}
	fmt.Println(p1.nombre)
	p1.edad = 242
	fmt.Println(p1.edad)

	//Punteros y Metodos
	var x int = 10
	var px *int = &x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(px)
	editar(&x)
	fmt.Println(x)
	//metodo
	pliss := Persona{"Lis Mesa ", 27, "SuCorreo"}
	pliss.diHola()

}

func editar(x *int) {
	*x = 45
}
