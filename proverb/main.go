package main

import "fmt"

func main() {
	fmt.Println(Proverb([]string{"nail", "shoe"}))
}

func Proverb(rhyme []string) (proverbSlice []string) {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if len(rhyme) == 0 {
		return proverbSlice
	}
	if len(rhyme) == 1 {
		return append(proverbSlice, fmt.Sprintf("And all for the want of a %s.", string(rhyme[0])))
	}
	proverb := "For want of a %s the %s was lost."
	for i := 0; i < len(rhyme)-1; i++ {
		word1 := rhyme[i]
		word2 := rhyme[i+1]
		proverbSlice = append(proverbSlice, fmt.Sprintf(proverb, word1, word2))
	}
	lastLine := fmt.Sprintf("And all for the want of a %s.", string(rhyme[0]))
	proverbSlice = append(proverbSlice, lastLine)
	return proverbSlice
}
