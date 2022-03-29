package main

import "fmt"

func main() {
	sum := 0
	a := [5]int{0, 1, 2, 3, 4}
	for _, value := range a {
		sum += value
	}
	fmt.Println(sum)
}
