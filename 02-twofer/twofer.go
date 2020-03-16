package twofer

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	s := "One for %s, one for me."
	if name != "" {
		s = fmt.Sprintf(s, name)
	} else {
		return fmt.Sprintf(s, "you")
	}
	return s
}
