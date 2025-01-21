package main

import "fmt"

func GD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func GD2(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	a := 10050
	b := 2555500
	fmt.Println(GD(a, b))
	fmt.Println(GD2(a, b))
}
