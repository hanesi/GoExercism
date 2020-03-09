package main

import "fmt"

type Planet string

func main() {
	ageCalc(1000000000)
}

func ageCalc(Age float64) {
	dict := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 64.79132,
	}

	a := Age / 31557600

	for k, v := range dict {
		dict[k] = v * a
	}
	fmt.Println(dict)
}
