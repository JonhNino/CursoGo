package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contanctos
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guarda Contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewDecoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// Carga contactos desde un archivo_json
func loadContactsFromfile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

}

func divide(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("NO puedes dividir por cero")
	}
}
func main() {
	/*str := "123g"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Numero Conver:", num)
	*/
	/*
		result, err := divide(10, 2)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Resultado", result)

		fmt.Println("..............................")
		defer fmt.Println(3)
		defer fmt.Println(1)
		defer fmt.Println(32)
		file, err := os.Create("hola.texr")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		_, err = file.Write([]byte("Hola, Jonh Niño A"))
		if err != nil {
			fmt.Println(err)
			return
		}
	*/

	divide(100, 10)
	divide(200, 25)
	divide(34, 0)
	divide(100, 4)

	///Registro de Errores
	//log.SetPrefix("main():")
	//log.Panic("Registro Panico")
	//log.Fatal("Registro Errores")
	//log.Print("Este es un mensaje de registro")
	//log.Println("Este es otro mensaje de registro")
	//log.SetPrefix("main():")
	//file, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0644)
	//if err != nil {
	//		log.Fatal(err)
	//}
	//log.SetOutput(file)
	//log.Println("Un Log")

	//Proyecto
	//Slice de contactos
	var contacts []Contact

	//Cargar contactos existentes desde el archivo
	err := loadContactsFromfile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}
	// Crear instancia de flujo
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("===GESTOR DE CONTACTOS ===\n",
			"1. Agregar un Contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opcion: ")

		//Leer una opcion
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		//Manejar Opciiones Usuario
		switch option {
		case 1:
			//Ingresar y crear Contact
			var c Contact
			fmt.Printf("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Printf("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Printf("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			//Agregar un nuevo Contacto a Slice
			contacts = append(contacts, c)
			///Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)

			}
		case 2:
			//Mostrar todos los contactos
			fmt.Println("=====================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre : %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
		case 3:
			//Salir del programa
			return
		default:
			fmt.Println("Opcion no Valida ")

		}
	}
}
