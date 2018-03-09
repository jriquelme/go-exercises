package main

// LargestProduct returns the index of the two positive integers in a list whose product is the largest.
//
// found will be false if a suitable pair of numbers wasn't found. The index of the largest number is returned in i0,
// the smaller in i1.
func LargestProduct(numbers ...int) (found bool, i0, i1 int) {
	i0 = -1
	i1 = -1
	for i, n := range numbers {
		if i0 == -1 || n > numbers[i0] {
			i1 = i0
			i0 = i
		} else if i1 == -1 || n > numbers[i1] {
			i1 = i
		}
	}
	found = i0 >= 0 && i1 >= 0 && numbers[i0] > 0 && numbers[i1] > 0
	return
}
