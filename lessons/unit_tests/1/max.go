package main

func Max(slice []int) int {
	var max int

	for _, value := range slice {
		if value > max {
			max = value
		}
	}

	return max
}

func MaxIndex(slice []int) int {
	var maxIndex int

	for i, value := range slice {
		if value > slice[maxIndex] {
			maxIndex = i
		}
	}

	return maxIndex
}
