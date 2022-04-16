package main

import "fmt"

func main() {
	fmt.Println(Calculate(2))
}

func Calculate(a int) int {
	res := a + 2

	return res
}
