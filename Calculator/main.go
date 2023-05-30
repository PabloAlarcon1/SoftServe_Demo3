package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Solicitar al usuario ingresar los dos números
	fmt.Println("Ingrese el primer número:")
	num1, err := obtenerNumero()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Ingrese el segundo número:")
	num2, err := obtenerNumero()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Mostrar las operaciones disponibles
	fmt.Println("Seleccione la operación:")
	fmt.Println("1. Suma")
	fmt.Println("2. Resta")
	fmt.Println("3. Multiplicación")
	fmt.Println("4. División")

	// Leer la opción seleccionada por el usuario
	var opcion int
	fmt.Scanln(&opcion)

	// Realizar la operación correspondiente y mostrar el resultado
	switch opcion {
	case 1:
		resultado := num1 + num2
		fmt.Println("Resultado:", resultado)
	case 2:
		resultado := num1 - num2
		fmt.Println("Resultado:", resultado)
	case 3:
		resultado := num1 * num2
		fmt.Println("Resultado:", resultado)
	case 4:
		if num2 == 0 {
			fmt.Println("Error: división por cero")
		} else {
			resultado := num1 / num2
			fmt.Println("Resultado:", resultado)
		}
	default:
		fmt.Println("Opción no válida")
	}
}

// Función para obtener un número ingresado por el usuario
func obtenerNumero() (float64, error) {
	var input string
	fmt.Scanln(&input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}