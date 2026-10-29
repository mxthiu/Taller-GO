package main

func sortAscending(values []int) []int {
	// Crea copia
	sorted := make([]int, len(values))

	for i := 0; i < len(values); i++ {
		sorted[i] = values[i]
	}

	// Bubble Sort
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-1-i; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	return sorted
}
