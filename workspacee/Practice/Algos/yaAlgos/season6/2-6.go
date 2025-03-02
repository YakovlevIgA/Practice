// https://contest.yandex.ru/contest/66793/problems/F/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const mod = 1000000007

func find(n int, nums []int) int { // 1

	prefSums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefSums[i] = (prefSums[i-1] + nums[i-1]) % mod
	}

	// Суффиксные суммы
	suffSums := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suffSums[i] = (suffSums[i+1] + nums[i]) % mod
	}

	// Подсчет ответа
	ans := 0
	for nowPos := 1; nowPos < n-1; nowPos++ {
		ans = (ans + prefSums[nowPos]*nums[nowPos]%mod*suffSums[nowPos+1]%mod) % mod
	}

	return ans
} // 1

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
