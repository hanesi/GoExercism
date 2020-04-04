/*
Convert map[int][]string{} to map[string]int{}
*/
package main

import "fmt"

func main() {
	oldMap := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}

	fmt.Println(Transform(oldMap))
}

// Transform takes an old format pf scrabble dict, changing it to a new one
func Transform(old map[int][]string) map[string]int {
	newMap := map[string]int{}
	for k, v := range old {
		for _, val := range v {
			newMap[val] = k
		}
	}
	return newMap
}
