/*
This function accepts an int and returns whether or not it's a leap year
*/
package main

import "fmt"

func main() {
	fmt.Println(IsLeapYear(1992))
}

// IsLeapYear should have a comment documenting it.
func IsLeapYear(num int) bool {
	fourCheck := num % 4
	fourHundCheck := num % 400
	hundCheck := num % 100

	if fourCheck == 0 && (fourHundCheck == 0 || hundCheck != 0) {
		return true
	}
	return false
}
