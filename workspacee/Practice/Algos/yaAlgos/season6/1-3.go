// https://contest.yandex.ru/contest/66792/enter/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Print("Введите количество строк n: ")
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Очищаем буфер, чтобы пропустить оставшийся символ новой строки

	lines := make([]string, n) // Сохраняем строки в срез

	fmt.Println("Введите строки:")

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			lines[i] = scanner.Text() // Считываем строку и сохраняем ее
		}
	}
	fmt.Println(lines)
}
