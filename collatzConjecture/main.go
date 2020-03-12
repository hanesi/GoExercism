package main

import "fmt"

func main() {
	fmt.Println(CollatzConjecture(0))
}

func CollatzConjecture(start int) (int, error) {
	if start <= 0 {
		return 0, fmt.Errorf("%d is not an allowed value", start)
	}
	var count int
	for count = 0; start > 1; count++ {
		switch {
		case start%2 == 0:
			start = start / 2
		case start%2 != 0:
			start = (3 * start) + 1
		}
	}
	return count, nil
}
