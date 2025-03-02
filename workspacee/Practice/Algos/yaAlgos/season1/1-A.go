// https://contest.yandex.ru/contest/27393/problems/A/

package main

import "fmt"

func temp(tRoom, tCond int, mode string) int {
	var ans int
	if mode == "freeze" {
		switch {
		case tRoom > tCond:
			ans = tCond
		case tRoom == tCond:
			ans = tCond
		case tRoom < tCond:
			ans = tRoom
		}
	}
	if mode == "heat" {
		switch {
		case tRoom > tCond:
			ans = tRoom
		case tRoom == tCond:
			ans = tCond
		case tRoom < tCond:
			ans = tCond
		}
	}
	if mode == "auto" {
		switch {
		case tRoom > tCond:
			ans = tCond
		case tRoom == tCond:
			ans = tCond
		case tRoom < tCond:
			ans = tCond
		}
	}

	if mode == "fan" {
		switch {
		case tRoom > tCond:
			ans = tRoom
		case tRoom == tCond:
			ans = tRoom
		case tRoom < tCond:
			ans = tRoom
		}

	}
	return ans
}

func main() {
	var tRoom int
	var tCond int
	fmt.Scanln(&tRoom, &tCond)
	var mode string
	fmt.Scan(&mode)
	fmt.Println(temp(tRoom, tCond, mode))

}
