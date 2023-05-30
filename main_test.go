package main

import (
	"testing"
)

func TestOperaciones(t *testing.T) {
	t.Run("Suma", func(t *testing.T) {
		num1 := 2.5
		num2 := 3.7
		expected := 6.2

		resultado := suma(num1, num2)
		if resultado != expected {
			t.Errorf("Suma incorrecta. Se esperaba %f pero se obtuvo %f", expected, resultado)
		}
	})

	t.Run("Resta", func(t *testing.T) {
		num1 := 5.2
		num2 := 3.1
		expected := 2.1

		resultado := resta(num1, num2)
		if resultado != expected {
			t.Errorf("Resta incorrecta. Se esperaba %f pero se obtuvo %f", expected, resultado)
		}
	})

	t.Run("Multiplicacion", func(t *testing.T) {
		num1 := 30.0
		num2 := 5.0
		expected := 150.0

		resultado := multiplicacion(num1, num2)
		if resultado != expected {
			t.Errorf("Resta incorrecta. Se esperaba %f pero se obtuvo %f", expected, resultado)
		}
	})

	t.Run("Division", func(t *testing.T) {
		num1 := 30.0
		num2 := 5.0
		expected := 6.0

		resultado := division(num1, num2)
		if resultado != expected {
			t.Errorf("Resta incorrecta. Se esperaba %f pero se obtuvo %f", expected, resultado)
		}
	})

}

func suma(a, b float64) float64 {
	return a + b
}

func resta(a, b float64) float64 {
	return a - b
}

func multiplicacion(a, b float64) float64 {
	return a * b
}

func division(a, b float64) float64 {
	return a / b
}
