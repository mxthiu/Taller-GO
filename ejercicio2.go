package main

import "fmt"

//Ejericio 2: Estadisticas de un arreglo

func suma(values []int) int {
	total := 0

	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	return total
}

func minimo(values []int) int {
	min := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min
}

func maximo(values []int) int {
	max := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}

	return max
}

func contarNegativos(values []int) int {
	contador := 0

	for i := 0; i < len(values); i++ {
		if values[i] < 0 {
			contador++
		}
	}

	return contador
}

func recorridoInverso(values []int) {
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Print(values[i], " ")
	}
	fmt.Println()
}
