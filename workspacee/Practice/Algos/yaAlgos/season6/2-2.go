// https://contest.yandex.ru/contest/66793/problems/B/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findAns(n, k int, numbers []int) int {
	sumK := 0
	left, right := 0, 0
	var inRow int

	for left < n {
		for right < n && (right == left || sumK <= k) {
			sumK += numbers[right]
			if sumK == k {
				inRow += 1
			}
			right += 1
		}
		sumK -= numbers[left]
		if sumK == k {
			inRow += 1
		}
		left += 1
	}
	return inRow

}

func main() {
	var n, k int
	_, err := fmt.Scanln(&n, &k)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка считывания", err)
		return
	}
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	var numbers []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка преобразования", err)
			return
		}
		numbers = append(numbers, num)
	}

	fmt.Println(findAns(n, k, numbers))

}
