package main

import (
	"fmt"
)

func main() {
	biologi := []string{"Didik Prabowo", "Charly Van Houten"}
	fisika := make([]string, len(biologi))

	copy(fisika, biologi)
	fmt.Println(fisika)
}
