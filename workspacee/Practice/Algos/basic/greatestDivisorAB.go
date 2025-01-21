// Евклид - наибольший общий делитель числа a и b.

package main

import "fmt"

func greatestDivisor1(a, b int) int { // Итеративное вычитание, сложность O(n), сравнения накинут
	// в районе 10% времени за каждое.
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func greatestDivisor2(a, b int) int { // Итеративное деление с остатком, сложность O(log(min(a,b)))
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {

	a := 451520158000
	b := 105500

	fmt.Printf("Итеративное вычитание, ответ: %v\n", greatestDivisor1(a, b))
	fmt.Printf("Итеративное деление с остатком, ответ: %v\n", greatestDivisor2(a, b))
}
