package main

// Ejercicio 7: Contar el agua atrapada
func buildLeftMax(height []int, leftMax []int) {
	if len(height) == 0 {
		return
	}

	leftMax[0] = height[0]

	for i := 1; i < len(height); i++ {
		if height[i] > leftMax[i-1] {
			leftMax[i] = height[i]
		} else {
			leftMax[i] = leftMax[i-1]
		}
	}
}

func buildRightMax(height []int, rightMax []int) {
	if len(height) == 0 {
		return
	}

	last := len(height) - 1
	rightMax[last] = height[last]

	for i := last - 1; i >= 0; i-- {
		if height[i] > rightMax[i+1] {
			rightMax[i] = height[i]
		} else {
			rightMax[i] = rightMax[i+1]
		}
	}
}

//Ejercicio 8: Trapping Rain Water
func trappedWater(height []int) int64 {
	if len(height) == 0 {
		return 0
	}

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	buildLeftMax(height, leftMax)
	buildRightMax(height, rightMax)

	total := int64(0)

	for i := 0; i < len(height); i++ {

		var limite int

		if leftMax[i] < rightMax[i] {
			limite = leftMax[i]
		} else {
			limite = rightMax[i]
		}

		agua := limite - height[i]

		total += int64(agua)
	}

	return total
}
