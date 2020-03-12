package main

import "fmt"

func main() {
	fmt.Println(CollatzConjecture(1000000))
}

func CollatzConjecture(start int) int {
	var count int
	for count = 0; start > 1; count++ {
		switch {
		case start%2 == 0:
			start = start / 2
		case start%2 != 0:
			start = (3 * start) + 1
		}
	}
	return count
}
