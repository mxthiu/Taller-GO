package main

// Ejercicio 3: copia independiente con make y un ciclo
func copyArray(values []int) []int {
	result := make([]int, len(values))
	for i := 0; i < len(values); i++ {
		result[i] = values[i]
	}
	return result
}
