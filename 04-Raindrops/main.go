/*
This function takes a number and converts it to a string of raindrop
noises based on it's factors which are defined in the strMap variable
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Convert(5))
}

// Convert number to raindrop sound
func Convert(factor int) string {
	strMap := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	returnStr := ""

	if factor%3 == 0 {
		returnStr += strMap[3]
	}
	if factor%5 == 0 {
		returnStr += strMap[5]
	}
	if factor%7 == 0 {
		returnStr += strMap[7]
	}

	if len(returnStr) == 0 {
		return strconv.Itoa(factor)
	}

	return returnStr
}
