package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(AddGigasecond(time.Now().Local()))
}

func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}
