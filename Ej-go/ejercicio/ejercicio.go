package main

import (
	"fmt"
	"os"
)

func main() {
	var path, text string
	path = "/home/lautaro/prueba"
	text = "hola mundo"
	f := create_file(path)
	defer closeFile(f)
	write_file(f, text)
}

func create_file(path string) *os.File {
	f, _ := os.Create(path)
	return f
}

func write_file(f *os.File, text string) {
	fmt.Fprintln(f, text)
}

func closeFile(f *os.File) {
	f.Close()
}
