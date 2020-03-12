/*
Function that takes a string and returns an acronym
Complementary metal-oxide semiconductor becomes "CMOS"
*/
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Abbreviate("Complementary metal-oxide semiconductor"))
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var buffer bytes.Buffer
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	stringList := strings.Split(s, " ")
	for _, v := range stringList {
		buffer.WriteByte(v[0])
	}
	return buffer.String()
}
