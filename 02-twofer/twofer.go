package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	var s string
	if name != "" {
		s = fmt.Sprintf("One for %s, one for me.", name)
	} else {
		return "One for you, one for me."
	}
	return s
}
