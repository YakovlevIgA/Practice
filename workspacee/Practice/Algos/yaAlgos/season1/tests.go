package main

import (
	"fmt"
	"strings"
)

func remove(input string) string {
	replacer := strings.NewReplacer("+", "", "-", "", "(", "", ")", "")
	return replacer.Replace(input)
}

func main() {

	num1 := "8(495)430-23-97"
	// num2 := "+7-4-9-5-43-023-97"
	//for i, num := range num1 {

	num1 = remove(num1)
	fmt.Println(num1)
	//}
}
