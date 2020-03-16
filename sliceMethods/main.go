package main

import (
	"fmt"
	"strings"
)

func main() {
	l1 := []string{"a", "b", "c"}
	l2 := []string{"d", "e", "f"}
	l3 := []string{"g", "h", "i"}

	mytest := func(s string) bool { return !strings.Contains(s, "a") && len(s) <= 7 }

	fmt.Println(concatLists(l1, l2, l3))
	fmt.Println(appendLists(l1, l2))
	fmt.Println(getLength(l1))
	fmt.Println(reverseList(l1))
	fmt.Println(filter(l1, mytest))
	fmt.Println(Map(l2, strings.ToUpper))
}

func concatLists(series ...[]string) []string {
	new := []string{}
	for _, v := range series {
		new = append(new, v...)
	}
	return new
}

func appendLists(l1 []string, l2 []string) []string {
	return append(l1, l2...)
}

func getLength(list []string) int {
	return len(list)
}

func reverseList(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
