// https://contest.yandex.ru/contest/66793/problems/D/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func find(n, k int, nums []int) int {
	sort.Ints(nums)
	left := 0
	right := 1
	var ans int
	if n == 1 {
		ans = 1
	} else {
		for left < n && right < n {
			if nums[right]-nums[left] <= k {
				if ans <= right-left+1 {
					ans = right - left + 1
				}

				right += 1
			} else {
				left += 1
			}
		}
	}
	return ans
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
	fmt.Println(find(n, k, numbers))

}
