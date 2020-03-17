/*
This function takes 3 floats and tells you what kind of
triangle it would form
*/
package main

import "fmt"

type Kind string

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

func main() {
	fmt.Println(KindFromSides(2.2, 2.2, 3.5))
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
		return k
	}
	if a > 0 && a == b && b == c {
		k = Equ
		return k
	}
	if a != b && b != c && a != c {
		k = Sca
		return k
	}
	k = Iso
	return k
}
