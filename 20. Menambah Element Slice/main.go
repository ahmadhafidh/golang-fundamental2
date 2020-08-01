package main

import (
	"fmt"
)

func main() {
	mahasiswalama := []string{"Didik Prabowo", "Charly Van Houten"}
	mahasiswabaru := "Udin Kumalasari"
	mahasiswa := append(mahasiswalama, mahasiswabaru)
	fmt.Println(mahasiswa)
}
