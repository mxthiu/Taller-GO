package main

import "fmt"

// Ejercicio 6: 3Sum sin duplicados, sobre una copia ordenada
func threeSum(values []int) int {
	n := len(values)
	if n < 3 {
		return 0
	}

	s := sortAscending(values)

	count := 0
	for i := 0; i < n-2; i++ {
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			sum := s[i] + s[left] + s[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				fmt.Printf("Terna: (%d, %d, %d)\n", s[i], s[left], s[right])
				count++
				left++
				right--
				// Saltar repetidos del lado izquierdo y derecho
				for left < right && s[left] == s[left-1] {
					left++
				}
				for left < right && s[right] == s[right+1] {
					right--
				}
			}
		}
	}
	return count
}
