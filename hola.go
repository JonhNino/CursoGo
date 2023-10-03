package main

import (
	"fmt"
	"strconv"

	"math"

	"rsc.io/quote"
)

// Declaracion Constantes
const Pi1 = 3.14
const (
	x = 100
	y = 0b1010 //Binario
	z = 0o12   //Octal
	w = 0xFF   //Hexa
)
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

const precision = 2 // Número de decimales de precisión

func main() {
	fmt.Println("Hola Munndo")
	fmt.Println(quote.Go())
	//Declaracion Variables

	/*
		var firstName, lastaName string
		var age int
	*/

	/*
		var (
			firstName = "Jonh"
			lastaName = "Nino"
			age       = 30
		)
	*/
	/*
		var (
			firstName, lastaName, age = "Jonh", "Nino", 30
		)
	*/

	firstName, lastaName := "Jonh", "Nino"
	age := 30

	fmt.Println(firstName, lastaName, age)
	fmt.Println(Pi1)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)
	fullName := "Jonh Niño A \t(Cargo \"Ing. Electronico\")\n"
	fmt.Println(fullName)
	var a byte = 'a'
	fmt.Println(a)
	s := "Hola"
	fmt.Println(s[0])
	var r rune = '😻'
	fmt.Println(r)
	var (
		defaultInit   int
		defaultUint   uint
		defaulFloat   float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println(defaultInit, defaultUint, defaulFloat, defaultBool, defaultString)

	//Conversiones de Tipos
	var integer16 int16 = 50
	var integer32 int32 = 100
	fmt.Println(integer16 + int16(integer32))

	f := "100"
	i, _ := strconv.Atoi(f)
	fmt.Println(i + i)
	n := 424
	f = strconv.Itoa(n)
	fmt.Println(f + f)
	var name, name2 string
	var edad int
	fmt.Println("Ingres un Nombre: ")
	fmt.Scanln(&name, &name2)
	fmt.Println("Ingrese la edad:")
	fmt.Scanln(&edad)
	/*name := "Liss"
	edad := 28
	*/

	fmt.Println("Otro Mensaje")
	fmt.Printf("Hola, Me llama %s %s y tengo %d años.  \n", name, name2, edad)
	greeting := fmt.Sprintf("Hola, Me llama %s y tengo %d años.  \n", name, edad)
	fmt.Print(greeting)

	//Ejercicio
	var lado1, lado2, hipotenusa, area, perimetro float64

	// Solicitar al usuario la longitud de los lados
	fmt.Print("Ingrese lado 1: ")
	fmt.Scan(&lado1)

	fmt.Print("Ingrese lado 2: ")
	fmt.Scan(&lado2)

	// Calcular la hipotenusa usando el teorema de Pitágoras
	hipotenusa = math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))

	// Calcular el área del triángulo
	area = (lado1 * lado2) / 2

	// Calcular el perímetro del triángulo
	perimetro = lado1 + lado2 + hipotenusa

	// Imprimir el área y el perímetro con dos decimales de precisión
	fmt.Printf("\nÁrea: %.*f\n", precision, area)
	fmt.Printf("Perímetro: %.*f\n", precision, perimetro)

}
