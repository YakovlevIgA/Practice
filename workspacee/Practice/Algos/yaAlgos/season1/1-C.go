// https://contest.yandex.ru/contest/27393/problems/C/
package main

import (
	"fmt"
	"strings"
)

func formating(number string) string {
	chars := "+()-"
	toRemove := make(map[rune]bool)
	for _, ch := range chars {
		toRemove[ch] = true
	}
	var result strings.Builder
	for _, ch := range number {
		if !toRemove[ch] {
			result.WriteRune(ch)
		}
	}
	format1 := result.String()
	var format2 string
	if len(format1) < 11 {
		format2 = "8495" + format1
	} else {
		format2 = format1
	}
	if string(format2[0]) == "7" {
		format2 = "8" + format2[1:]
	}
	return format2
}

func main() {
	var checkNumber, numbers1, numbers2, numbers3 string
	fmt.Scan(&checkNumber)
	fmt.Scan(&numbers1)
	fmt.Scan(&numbers2)
	fmt.Scan(&numbers3)
	checkNumber = formating(checkNumber)
	numbers1 = formating(numbers1)
	numbers2 = formating(numbers2)
	numbers3 = formating(numbers3)
	if checkNumber == numbers1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	if checkNumber == numbers2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	if checkNumber == numbers3 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
