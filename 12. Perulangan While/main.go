package main

import (
	"fmt"
)

func main() {
	benar := true
	i := 0

	for benar {
		if i == 3 {
			benar = false
		}
		fmt.Println("Angka ke", i)
		i++
	}
}
