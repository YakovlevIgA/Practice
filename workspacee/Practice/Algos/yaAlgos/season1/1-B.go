// https://contest.yandex.ru/contest/27393/problems/B/
package main

import "fmt"

func isReal(a, b, c int) string {
	if a+b > c && a+c > b && b+c > a {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	var a, b, c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	fmt.Println(isReal(a, b, c))

}
