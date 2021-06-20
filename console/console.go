package console

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	switch runtime.GOOS {
	case "darwin":
	case "linux":
		{
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			err := cmd.Run()

			if err != nil {
				fmt.Printf("Failed to clean screen, command finished with error %v", err)
			}
		}

	case "windows":
		{
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			err := cmd.Run()

			if err != nil {
				fmt.Printf("Failed to clean screen, command finished with error %v", err)
			}
		}

	default:
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func RequestString(requestMessage string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(requestMessage)
		result, err := reader.ReadString('\n')

		if err == nil {
			return result
		}

		fmt.Println("Ha ocurrido un fallo al leer la entrada de datos")
	}
}

func RequestUInt8(requestMessage string) uint8 {
	for {
		var input int
		fmt.Println(requestMessage)

		_, err := fmt.Scanf("%d", &input)

		if err == nil {
			return input
		}

		fmt.Println("Ha ocurrido un fallo al leer la entrada de datos")
	}
}

func RequestRune(requestMessage string) rune {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(requestMessage)
		result, _, err := reader.ReadRune()

		if err == nil {
			return result
		}

		fmt.Println("Ha ocurrido un fallo al leer la entrada de datos")
	}
}
