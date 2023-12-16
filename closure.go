package main

import "fmt"

func main() {
	x := 0
	numero := func() int {
		x += 5
		return x
	}
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
	fmt.Println(numero())
}
