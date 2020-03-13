package main

import "fmt"

func main() {
	l1 := []string{"a", "b", "c"}
	l2 := []string{"d", "e", "f"}
	l3 := []string{"g", "h", "i"}

	fmt.Println(concatLists(l1, l2, l3))
}

func concatLists(series ...[]string) []string {
	new := []string{}
	for _, v := range series {
		new = append(new, v...)
	}
	return new
}
