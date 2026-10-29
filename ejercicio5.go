package main

import "fmt"

// Ejercicio 5: dos valores que suman un objetivo.
func findPair(values []int, target int) bool {
	left := 0
	right := len(values) - 1
	for left < right {
		sum := values[left] + values[right]
		if sum == target {
			fmt.Printf("Un par válido es (%d, %d)\n", values[left], values[right])
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

func probarPar(values []int, target int) {
	if !findPair(values, target) {
		fmt.Println("No existe")
	}
}
