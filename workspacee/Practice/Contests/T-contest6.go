package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Считываем количество процессов
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	// Время выполнения процессов
	time := make([]int64, n+1)
	// Граф зависимостей
	adj := make([][]int, n+1)
	// Степень входа
	inDegree := make([]int, n+1)
	// Очередь для топологической сортировки
	queue := make([]int, 0)

	// Чтение входных данных и построение графа
	for i := 1; i <= n; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		ti, _ := strconv.ParseInt(parts[0], 10, 64)
		time[i] = ti
		for _, p := range parts[1:] {
			dep, _ := strconv.Atoi(p)
			adj[dep] = append(adj[dep], i)
			inDegree[i]++
		}
	}

	// Инициализация очереди с узлами степени входа 0
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Время завершения процессов
	completionTime := make([]int64, n+1)

	// Топологическая сортировка и вычисление времени завершения
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, next := range adj[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
			// Вычисляем время завершения процесса
			completionTime[next] = max(completionTime[next], completionTime[curr]+time[next])
		}
	}

	// Найти максимальное время завершения
	maxTime := int64(0)
	for i := 1; i <= n; i++ {
		maxTime = max(maxTime, completionTime[i])
	}

	fmt.Println(maxTime)
}

// Функция для нахождения максимального из двух значений
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
