package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	path = "/home/lautaro/ejemplo"
	f, _ := os.Create(path)
	fmt.Fprintln(f, "data")
	f.Close()
}
