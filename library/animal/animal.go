package animal

import "fmt"

type Animal interface {
	Sonido()
}

// Struct Prro y Methods
type Perro struct {
	Nombre string
}

func (p Perro) Sonido() {
	fmt.Println(p.Nombre + "Hace como Prro")
}

// Struct Gato y Methods
type Gato struct {
	Nombre string
}

func (p Gato) Sonido() {
	fmt.Println(p.Nombre + "Hace como Gato")
}
