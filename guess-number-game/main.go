package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	//If Else en Go
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("De Mañana IF")
	} else if t := time.Now(); t.Hour() < 17 {
		fmt.Println("Es Tarde IF")
	} else {
		fmt.Println("De Noche IF")
	}

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("De Mañana Switch")
	case t.Hour() < 17:
		fmt.Println("De Tarde Switch")
	default:
		fmt.Println("De Noche Switch")

	}

	//Switch Case en Go no es necesario el Break
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go Run -> Windows")
	case "linux":
		fmt.Println("Go Run -> linux")
	case "darwin":
		fmt.Println("Go Run -> darwin")
	default:
		fmt.Println("Go Run -> Otro OS")
	}

	/// Bucle For en Go
	// Bucle Infinito-Bucle tipico
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Bucle Tipico
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Continue  se utiliza para saltar al siguiente iterador
	//& Break=Salir antes de que la condicion de finalizacion se alcance
	for i := 1; i <= 10; i++ {
		fmt.Println(i)

		if i == 3 {
			continue

		}
		if i == 5 {
			break

		}
		fmt.Println(i)
	}
	// Uso de Funciones con Go
	hello("Jonh")
	saludo := hello1("Jonh")
	fmt.Println(saludo)
	sum, mul := calc(3, 5)
	fmt.Println("La Suma Es :", sum)
	fmt.Println("La Multiplicaion Es :", mul)
	sum1, mul1 := calc1(3, 5)
	fmt.Println("La Suma Es :", sum1)
	fmt.Println("La Multiplicaion Es :", mul1)

	/// Proyecto De Modulo
	fmt.Println("Inicio Proyecto")
	fmt.Println(rand.Intn(100))
	jugar()
}

// Funciones con Go
func hello(name string) {
	fmt.Println("Desde la Funcion", name)
}

func hello1(name string) string {
	return "Desde la Funcion que retorna " + name
}

func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

func calc1(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return sum, mul
}

// Funciones Proyecto
func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntents = 10
	for intentos < maxIntents {
		intentos++
		fmt.Printf("Ingrese un Numero (intentos restantes: %d)", maxIntents-intentos+1)
		fmt.Scanln(&numIngresado)
		if numIngresado == numAleatorio {
			fmt.Println("¡Felicidades Adivinaste el Numero!")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El Numero a adivinar es Mayor.")
		} else {
			fmt.Println("El Numero a adivinar es Menor.")
		}
	}

	fmt.Println("Se acabaron los intentos el numero era", numAleatorio)
	jugarNuevamente()
}
func jugarNuevamente() {
	var eleccion string
	fmt.Println("¿Quires Volver A Jugar? (s/n)")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("¡Gracias Por Jugar!")
	default:
		fmt.Println("ELeccion Invalidad. Intentalo Nuevamente.")
		jugarNuevamente()
	}
}
