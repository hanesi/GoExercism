/*
Function takes a slice of strings and an operation and returns
a slice that has the operation applied to it.
*/
package main

import (
	"fmt"
	"strings"
)

type Operation func(string) string

func main() {
	sl := []string{"one", "two", "three"}
	fmt.Println(Accumulate(sl, strings.ToUpper))
}

func Accumulate(s []string, f Operation) []string {
	results := []string{}
	for _, v := range s {
		results = append(results, f(v))
	}
	return results
}
