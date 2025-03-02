// https://contest.yandex.ru/contest/27393/problems/D/

package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	if (a > c*c-b) && (c*c-b != 0) || c < 0 {
		fmt.Println("NO SOLUTION")
		os.Exit(0)
	}
	if b == c && a == 0 {
		fmt.Println("MANY SOLUTIONS")
		os.Exit(0)
	}
	if (c*c-b)%a == 0 {
		fmt.Println((c*c - b) / a)
	} else {
		fmt.Println("NO SOLUTION")
	}

}
