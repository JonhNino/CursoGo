package grettings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Funcion Hello a persona especifica
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Nombre Vacio")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Varios saludos
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {

		}
		messages[name] = message
	}
	return messages, nil

}

// Saludos aleatorios
func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Que bueno verte, %v",
		"¡Saludo, %v!¡Encantadao de concocertte !",
	}
	return formats[rand.Intn(len(formats))]
}
