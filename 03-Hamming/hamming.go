package hamming

import "fmt"

// Distance between two DNA strings
func Distance(a, b string) (int, error) {
	hammingDistance := 0
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, fmt.Errorf("Strings need to be the same length")
	}
	for i := range ar {
		if ar[i] != br[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
