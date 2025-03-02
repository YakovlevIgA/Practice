//https://contest.yandex.ru/contest/66793/problems/E/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// func findAns(n int, nums []int) string {                                                             2.085s ver
//	var mid int
//	var ans string
//	for n >= 1 {
//		switch {
//		case n == 1:
//			ans += strconv.Itoa(nums[0])
//		case n%2 == 0:
//			mid = n/2 - 1
//			ans += strconv.Itoa(nums[mid]) + " "
//			nums = append(nums[:mid], nums[mid+1:]...)
//
//		case n%2 != 0:
//			mid = n / 2
//			ans += strconv.Itoa(nums[mid]) + " "
//			nums = append(nums[:mid], nums[mid+1:]...)
//
//		}
//		n -= 1
//	}
//	return ans
// }

func findAns(n int, nums []int) []int {
	result := make([]int, 0, 1)
	var left, right int
	switch {

	case n%2 == 0:
		left = n/2 - 1
		right = n / 2
		for left >= 0 || right < n {
			if left >= 0 {
				result = append(result, nums[left])
				left--
			}
			if right < n {
				result = append(result, nums[right])
				right++
			}
		}
	case n%2 != 0:
		mid := n / 2
		left = n/2 - 1
		right = n/2 + 1
		result = append(result, nums[mid])
		for left >= 0 || right < n {
			if left >= 0 {
				result = append(result, nums[left])
				left--
			}
			if right < n {
				result = append(result, nums[right])
				right++
			}
		}
	}
	return result
}

func intsToString(nums []int) string {
	var sb strings.Builder
	for i, num := range nums {
		if i > 0 {
			sb.WriteString(" ") // Добавляем пробел между числами
		}
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
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
	sort.Ints(nums)
	fmt.Println(intsToString(findAns(n, nums)))
}
