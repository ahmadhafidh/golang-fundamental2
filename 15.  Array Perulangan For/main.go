package main

import (
	"fmt"
)

func main() {
	var name [5]string

	name[0] = "Didik Prabowo"
	name[1] = "Prabowo Didik"
	name[3] = "Mas Didik Prabowo"
	name[4] = "Prabowo Golang"

	for i, v := range name {
		fmt.Println("Index ke ", i, "=>", v)
	}
}