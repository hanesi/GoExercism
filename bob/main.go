/*
Bob has simple answers to remarks:
Bob answers 'Sure.' if you ask him a question, such as "How are you?".
He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying anything.
He answers 'Whatever.' to anything else.
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Hey("TEST STRING?"))
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	question := strings.Contains(remark, "?")
	isYell := IsUpper(remark)

	switch {
	case question == true && isYell == true:
		return "Calm down, I know what I'm doing!"
	case question == false && isYell == true:
		return "Whoa, chill out!"
	case question == true && isYell == false:
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
