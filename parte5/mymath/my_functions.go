package mymath

// Receives an array of float numbers and returns its sum.
func Sum(elements []float64) float64 {
	sum := 0.0
	for i := 0; i < len(elements); i++ {
		sum += elements[i]
	}

	return sum
}

// Receives the amount of elements of the desired arithmetic progression (AP)
// and the difference between elements and returns a slice containing the
// desired AP.
func ArithmeticProgression(amount, difference int64) []int64 {
	var elements []int64

	number := int64(1)
	count := int64(1)
	for count <= amount {
		elements = append(elements, number)
		count++
		number += difference
	}

	return elements
}
