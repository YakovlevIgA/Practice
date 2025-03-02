package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find(n int, nums []int) int {
	sum := 0
	var mid int
	var ans int
	for _, num := range nums {
		sum += num
	}
	switch {
	case sum%2 == 0:
		mid = sum / 2
	case sum%2 != 0:
		mid = sum/2 + 1
	}
	fmt.Println(mid)
	countMid := 0
	number := 0
	for countMid < mid {
		countMid += nums[number]
		number += 1
	}
	number -= 1
	for i := 0; i < number; i++ {
		ans += nums[i] * (number - i)
	}
	counting := 1
	for i := number + 1; i < n; i++ {
		ans += nums[i] * counting
		counting++
	}
	return ans
}

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка считывания", err)
		return
	}
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	var nums []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка преобразования", err)
			return
		}
		nums = append(nums, num)
	}

	fmt.Println(find(n, nums))

}
