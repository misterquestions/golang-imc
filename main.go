package main

import (
	"bufio"
	"fmt"
	"golang-imc/console"
	"golang-imc/person"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var registeredPersons []person.Person

	go displayMenu(&registeredPersons)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func displayMenu(personRegister *[]person.Person) {
	var exitApp = false

	for {
		// Display menu options
		console.Clear()
		fmt.Println("1. Registrar usuario")
		fmt.Println("2. Lista de usuarios")

		// Handle actions
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println("Error al leer el caracter, ", err)
		}

		switch char {
		case '1':
			registerPerson(personRegister, reader)
			break
		default:
			fmt.Println("Opción inválida")
			break
		}

		if exitApp {
			break
		}
	}
}

func registerPerson(personRegister *[]person.Person, reader *bufio.Reader) {
	console.Clear()
	name := console.RequestString("Ingrese el nombre del usuario")
	age := console.RequestUInt8("Ingrese la edad del usuario")
	genderRune := console.RequestRune("Ingrese el género del usuario (H: Hombre, M: Mujer).")

	var gender person.Gender

	if genderRune == 'H' {
		gender = person.Male
	} else {
		gender = person.Female
	}

	insertedUser := person.NewPerson(name, age, gender)
	*personRegister = append(*personRegister, *insertedUser)
}
