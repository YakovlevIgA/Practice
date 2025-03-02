// Условия

// Каждый раз, возвращаясь с работы, Василий вынужден отстоять большую очередь, чтобы сесть в маршрутку.
// В очереди перед Василием находится n человек.
// Когда человек заходит в маршрутку, он идет к месту, находящемуся как можно ближе к середине.
// Если таких мест несколько, человек идет к месту с меньшим номером.
// Например, если в маршрутке всего 6 мест и места с номерами 3, 4 заняты,
// то следующий пассажир сядет на место под номером 2.
// Когда в маршрутке кончаются места, она уезжает, а на ее место приезжает новая.
// Василий по очереди записывал, на какие места садятся пассажиры. Что записал Василий?

// Формат входных данных

// Единственная строка содержит числа n (1 ⩽ n ⩽ 5 ∗ 10^5) и m (1 ⩽ m ⩽ 5 ∗ 10^5)  — количество
// пассажиров, стоящих в очереди, и вместимость маршрутки соответственно.

// Формат выходных данных

// В  i−й строке выведите, на какое место сел i−й пассажир.

package main

import (
	"fmt"
	"math"
)

var places []int

func placing(m int) []int { // рассадка
	places = make([]int, m+1)
	middle := (math.Ceil(float64(m) / 2)) // номер середины автобуса
	places[1] = int(middle)
	way := m % 2

	switch {
	case way > 0:
		for i := 2; i <= m; i += 2 { // ниже середины
			places[i] = int(middle) - i/2
		}
		for i := 3; i <= m; i += 2 { // выше середины
			places[i] = int(middle + (float64(i)-1)/2)
		}
	case way == 0:
		for i := 2; i <= m; i += 2 {
			places[i] = int(middle) + i/2
		}
		for i := 3; i <= m; i += 2 {
			places[i] = int(middle - (float64(i)-1)/2)
		}

	}
	return places[1:]
}

func main() {
	var n, m int // человек в очереди и вместимость маршрутки
	fmt.Println("Введите количество человек в очереди и вместимость маршрутки")
	fmt.Scan(&n, &m)
	num := int(math.Ceil(float64(n) / float64(m)))
	slice := make([]int, 0)

	for j := 0; j < num; j++ {
		if j == num-1 {
			a := n % m // остаток пассажиров в последнем автобусе
			slice = append(slice, placing(m)[:(a)]...)

		} else {
			slice = append(slice, placing(m)...)
		}
	}
	for _, i := range slice {
		fmt.Println(i)
	}

}
