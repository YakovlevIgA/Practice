// https://contest.yandex.ru/contest/66793/problems/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func prefSum(n int, numbers []int) []int {
	res := make([]int, 1)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + numbers[i]
		res = append(res, sum)
	}

	return res[1:]
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	strN, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка считывания n", err)
		return
	}
	inputN := strings.TrimSpace(strN)
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

	n, err := strconv.Atoi(inputN)

	preResult := prefSum(n, numbers)
	var result string
	for _, number := range preResult {
		strNumber := strconv.Itoa(number)
		result += strNumber + " "
	}
	fmt.Println(result)
}
