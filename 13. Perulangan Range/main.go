package main

import (
	"fmt"
)

func main() {

	data := []string{"didik", "Prabowo"}

	for index, v := range data {
		fmt.Println("Perulangan ke", index, v)
	}
}
