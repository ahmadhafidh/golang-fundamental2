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

	name[len(name)-1] = ""

	for index := 0; index < len(name); index++ {
		fmt.Println("Index ke ", index, "=>", name[index])
	}
}