package main

import (
	"fmt"
	"log"

	"github.com/jonhNino/CursoGo/grettings"
)

func main() {
	log.SetPrefix("grettings")
	log.SetFlags(0)
	names := []string{"Liss", "Juna", "Maria"}
	messages, error := grettings.Hellos(names)

	//message, error := grettings.Hello("")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(messages)
}
