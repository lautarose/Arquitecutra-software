package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	path = "/home/lautaro/ejemplo"
	f, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error al leer archivo")
		return
	}
	fmt.Print(string(f))
}
