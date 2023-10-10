package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {

	//	myBook := book.NewBook("The Go Programming Language", "Heman Weed", 300)
	/*
		var myBook = book.Book{
			Title:  "The Go Programming Language",
			Author: "Heman Weed",
			Pages:  300,
		}
	*/
	myBook := book.NewBook("The Go Programming Language 1", "Heman Weed 1", 300)
	myBook.PrintInfo()
	myBook.SetTitle("The Go Programming Language 1 Special Edition")
	fmt.Println(myBook.GetTitle())
	myBook.PrintInfo()

	//Composicion
	myTextBook := book.NewTextBook("Comunicacion", "Jonh", 324, "Santillana SAC", "Secundaria")
	myTextBook.PrintInfo()

	//Polimorfismo
	fmt.Println("**********************************************************")
	//myBook.PrintInfo()
	//myTextBook.PrintInfo()
	book.Print(myBook)
	book.Print(myTextBook)

	///Interface
	miPerro := animal.Perro{Nombre: "Max"}
	miGato := animal.Gato{Nombre: "Cartuline"}

	miPerro.Sonido()
	miGato.Sonido()

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Toby"},
		&animal.Gato{Nombre: "Pacho"},
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Cartuline"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
