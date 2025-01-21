// Ввод: количество чисел и сами числа.
// Вывод: Максимальные числа, которые получатся в результате удаления из них одной цифры.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Создаем буфер для ввода
	reader := bufio.NewReader(os.Stdin)

	// Считываем первую строку, которая содержит число n

	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(input))

	// Создаем срез для хранения строк
	lines := make([]string, n)

	// Считываем n строк построчно

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		lines[i] = strings.TrimSpace(line) // Убираем лишние пробелы и переносы
	}
	var number string
	answers := make([]string, n)
	for i := 0; i < n; i++ {
		number = lines[i]
		if len(number) == 1 {
			answers[i] = "0"
		} else {
			for k := 0; k < len(number)-1; k++ {
				if number[k] < number[k+1] {
					answers[i] = number[:k] + number[k+1:]
					break
				} else {
					answers[i] = number[:len(number)-1]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(answers[i])
	}
}
