// https://contest.yandex.ru/contest/66793/problems/I/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func find(n int, interest, profit, mood []int) string {

//	return
// }

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)

	input1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка считывания", err)
		return
	}
	input2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка считывания", err)
		return
	}
	input3, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка считывания", err)
		return
	}

	input1 = strings.TrimSpace(input1)
	input2 = strings.TrimSpace(input2)
	input3 = strings.TrimSpace(input3)

	parts1 := strings.Split(input1, " ")
	parts2 := strings.Split(input2, " ")
	parts3 := strings.Split(input3, " ")

	var interest []int
	var profit []int
	var mood []int

	for _, part := range parts1 {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка преобразования", err)
			return
		}
		interest = append(interest, num)
	}

	for _, part := range parts2 {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка преобразования", err)
			return
		}
		profit = append(profit, num)
	}

	for _, part := range parts3 {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Ошибка преобразования", err)
			return
		}
		mood = append(mood, num)
	}

	fmt.Println(interest, profit, mood)
}
