// https://contest.yandex.ru/contest/66793/problems/C/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPairs(n, r int, nums []int) int {
	ans := 0
	left := 0
	right := 1
	for left < n-1 {
		for right < n && nums[right]-nums[left] <= r {
			right += 1
		}
		ans += n - right
		left += 1
	}
	return ans
}

func main() {
	var n, r int
	_, err := fmt.Scanln(&n, &r)
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

	fmt.Println(findPairs(n, r, numbers))
}
