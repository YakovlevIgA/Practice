package main

import (
	"fmt"
)

func findPassword(s, required string, k int) string {
	// Преобразуем набор символов в множество для быстрого поиска
	requiredSet := make(map[rune]struct{})
	for _, ch := range required {
		requiredSet[ch] = struct{}{}
	}

	n := len(s)
	requiredCount := len(requiredSet)

	// Словарь для подсчета количества символов из набора в текущем окне
	windowCount := make(map[rune]int)

	left, right := 0, 0
	uniqueCount := 0
	bestStart := -1
	bestLength := 0

	for right < n {
		// Добавляем новый символ в окно
		if _, found := requiredSet[rune(s[right])]; found {
			windowCount[rune(s[right])]++
			if windowCount[rune(s[right])] == 1 {
				uniqueCount++
			}
		}
		right++

		// Уменьшаем окно до тех пор, пока оно по-прежнему содержит все требуемые символы
		for uniqueCount == requiredCount && (right-left) > k {
			if _, found := requiredSet[rune(s[left])]; found {
				windowCount[rune(s[left])]--
				if windowCount[rune(s[left])] == 0 {
					uniqueCount--
				}
			}
			left++
		}

		// Проверяем, есть ли подходящая подстрока
		if uniqueCount == requiredCount && (right-left) <= k {
			if bestStart == -1 || left > bestStart || (left == bestStart && (right-left) > bestLength) {
				bestStart = left
				bestLength = right - left
			}
		}
	}

	if bestStart == -1 {
		return "-1"
	}

	return s[bestStart : bestStart+bestLength]
}

func main() {
	var s, required string
	var k int

	// Ввод данных
	fmt.Scan(&s)
	fmt.Scan(&required)
	fmt.Scan(&k)

	// Находим возможный пароль
	result := findPassword(s, required, k)
	fmt.Println(result)
}
