package main

import (
	"fmt"
	"strings"
)

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

func main() {
	Scale("f#", "MmMMmMM")
}

// Scale takes in a tonic (starting note) and interval (ex: "MMmmMm")
func Scale(tonic string, interval string) {
	notes := len(interval)
	result := []string{}

	if len(tonic) == 2 && strings.Contains(tonic, "#") == true {
		start := 0
		for i, v := range sharps {
			if v == tonic {
				start = i
			}
		}
		try := append(sharps[start:], sharps[:start]...)
		fmt.Println(try)
	}

	fmt.Println(len(tonic))
	fmt.Println(strings.Contains(tonic, "#"))
	fmt.Println(result)
	fmt.Println(notes)
	fmt.Println(flats)
}
