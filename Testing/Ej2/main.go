package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 2
	b := 2
	res, _ := Division(a, b)
	fmt.Println(a, "/", b, "= ", res)
}

func Division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("no puedo dividir por 0")
	}
	return a / b, nil
}
