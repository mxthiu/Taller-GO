package main

import "fmt"

func main() {
	values := []int{4, -2, 0, 7, -5}
	fmt.Println("Entrada:", values)

	// Ejercicio 1

	fmt.Println("--Ejercicio 1---")
	a, b := 8, -3
	fmt.Println("Ejercicio 1 antes:", a, b)
	swap(&a, &b)
	fmt.Println("Ejercicio 1 después:", a, b)

	// Ejercicio 2
	fmt.Println("--Ejercicio 2---")
	fmt.Println("Ejercicio 2, suma:", suma(values))
	fmt.Println("Ejercicio 2, mínimo:", minimo(values))
	fmt.Println("Ejercicio 2, máximo:", maximo(values))
	fmt.Println("Ejercicio 2, negativos:", contarNegativos(values))
	fmt.Println("Ejercicio 2, recorrido inverso:")
	recorridoInverso(values)

	// Ejercicio 3
	fmt.Println("--Ejercicio 3---")
	original := []int{4, -2, 0, 7, -5}
	cp := copyArray(original)
	cp[0] = 99
	fmt.Println("Ejercicio 3 original:", original)
	fmt.Println("Ejercicio 3 copia:", cp)

	// Ejercicio 4
	fmt.Println("--Ejercicio 4---")
	ordenado := sortAscending(values)

	fmt.Println("Ejercicio 4 original:", values)
	fmt.Println("Ejercicio 4 ordenado:", ordenado)

	// Ejercicio 5
	fmt.Println("--Ejercicio 5---")
	fmt.Println("Ejercicio 5, objetivo 1:")
	probarPar([]int{-4, -1, 0, 2, 5}, 1)
	fmt.Println("Ejercicio 5, objetivo 8:")
	probarPar([]int{-4, -1, 0, 2, 5}, 8)
	fmt.Println("Ejercicio 5, [0, 0] objetivo 0:")
	probarPar([]int{0, 0}, 0)
	fmt.Println("Ejercicio 5, [3] objetivo 6:")
	probarPar([]int{3}, 6)

	// Ejercicio 6
	fmt.Println("--Ejercicio 6---")
	fmt.Println("3Sum [-1, 0, 1, 2, -1, -4]:", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println("3Sum [0, 1, 1]:", threeSum([]int{0, 1, 1}))
	fmt.Println("3Sum [0, 0, 0]:", threeSum([]int{0, 0, 0}))
	fmt.Println("3Sum [0, 0, 0, 0]:", threeSum([]int{0, 0, 0, 0}))
	fmt.Println("3Sum [-2, 0, 0, 2, 2]:", threeSum([]int{-2, 0, 0, 2, 2}))

	// Ejercicio 7
	fmt.Println("--Ejercicio 7---")
	height := []int{3, 0, 2, 0, 4}

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	fmt.Println("Ejercicio 7")
	fmt.Println("height   :", height)
	fmt.Println("leftMax  :", leftMax)
	fmt.Println("rightMax :", rightMax)

	// Ejercicio 8
	fmt.Println("--Ejercicio 8---")
	fmt.Println("Ejercicio 8")

	fmt.Println("[5] ->",
		trappedWater([]int{5}),
	)

	fmt.Println("[0,1,2,3] ->",
		trappedWater([]int{0, 1, 2, 3}),
	)

	fmt.Println("[3,2,1,0] ->",
		trappedWater([]int{3, 2, 1, 0}),
	)

	fmt.Println(
		"[2,2,2] ->",
		trappedWater([]int{2, 2, 2}),
	)

	fmt.Println("[3,0,3] ->",
		trappedWater([]int{3, 0, 3}),
	)

	fmt.Println("[3,0,2,0,4] ->",
		trappedWater([]int{3, 0, 2, 0, 4}),
	)

	fmt.Println("[4,2,0,3,2,5] ->",
		trappedWater([]int{4, 2, 0, 3, 2, 5}),
	)

	fmt.Println("[0,1,0,2,1,0,1,3,2,1,2,1] ->",
		trappedWater([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}),
	)

	// caso propio de prueba
	fmt.Println(" caso propio [2,0,2] ->", trappedWater([]int{2, 0, 2}))
}
