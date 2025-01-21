// Есть массив длиной канатов, например, [9, 7, 5, 3]. Необходимо нарезать их на куски одинаковой длины так, чтобы получилось, скажем, 4 куска.
// Нужно найти максимальную возможную длину куска, при которой можно нарезать хотя бы 4 куска.

package main

import (
	"fmt"
)

func leftBinarySearch(arr []int, target int) int { // поиск первого элемента удовлетворяющего условие
	low, high := 0, len(arr)
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] >= target {
			high = mid // пытаемся найти первый элемент, удовлетворяющий условию
		} else {
			low = mid + 1
		}
	}
	return low // возвращаем индекс первого элемента, который >= target
}

func rightBinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)
	for low < high {
		mid := low + (high-low)/2
		if arr[mid] > target {
			high = mid // уменьшаем диапазон, если элемент больше target
		} else {
			low = mid + 1 // если меньше или равен target, смещаем нижнюю границу
		}
	}
	return low - 1 // возвращаем индекс последнего элемента, который <= target
}

func main() {
	arr := []int{1, 3, 3, 5, 7, 9}
	target := 3
	indexR := rightBinarySearch(arr, target)
	indexL := leftBinarySearch(arr, target)
	fmt.Println(indexL, indexR)

}
