/*
This code receives a DNA string and returns an equivalent RNA string
*/
package main

import "fmt"

func main() {
	fmt.Println(ToRNA("GATACA"))
}

func ToRNA(dna string) string {
	rna := ""
	mapper := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	for _, v := range dna {
		rna += mapper[string(v)]
	}
	return rna
}
