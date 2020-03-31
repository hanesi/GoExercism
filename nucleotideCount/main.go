package main

import (
	"errors"
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

func main() {
	var test DNA
	test = "GGGGGGG"
	fmt.Println(test.Counts())
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'G': 0,
		'A': 0,
		'T': 0,
		'C': 0,
	}

	for _, nucleotide := range d {
		if _, ok := h[nucleotide]; ok {
			h[nucleotide]++
		} else {
			return h, errors.New("invalid nucleotide found")
		}
	}
	return h, nil
}
