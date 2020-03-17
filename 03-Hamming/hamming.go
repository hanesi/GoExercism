package hamming

import "fmt"

// Calculate the hamming distance distance between two DNA string
func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("Strings need to be the same length")
	}
	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
