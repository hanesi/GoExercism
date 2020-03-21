package main

import "math"
import "fmt"

func main() {
	fmt.Println(darts(0, -1))
}

func darts(x, y float64) float64 {
	distance := math.Sqrt((math.Pow((math.Abs(x)), 2)) + (math.Pow((math.Abs(y)), 2)))
	switch {
	case distance > 10:
		return 0
	case distance <= 10 && distance > 5:
		return 1
	case distance <= 5 && distance > 1:
		return 5
	}
	return 10
}
