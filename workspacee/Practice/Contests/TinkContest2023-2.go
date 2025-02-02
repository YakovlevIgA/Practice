// Василий — фермер в двумерном государстве под названием Флатландия.
// У Василия есть 4n единиц забора, являющихся отрезками длины 1.
// Единицы забора разрешается располагать только параллельно осям Ox и Oy,
// при этом крайние точки забора должны оказаться в целочисленных точках.

// Василий хочет построить ровно k загонов для овечек, имеющих размеры 1 × 1.
// Каждый загон обязан иметь форму квадрата.
// Какое минимальное количество овечек могут вмещать k загонов, построенных Василием?
// Какое максимальное количество овечек могут вмещать k загонов, построенных Василием?

// Формат входных данных
// Первая строка содержит число n — количество четверок однометровых кусочков забора,
// которыми располагает Василий.
// Вторая строка содержит число k — количество загонов, которое хочет построить Василий.
// Гарантируется, что выполнена цепочка неравенств 1 ≤ k ≤ n ≤ 10^9.

// Формат выходных данных
// Через пробел выведите минимальную и максимальную суммарные площади загонов,
// которые могут иметь k загонов, сооруженных Василием.

// Замечания
// В первом примере для минимизации размеров, необходимо сделать размеры загонов 1×1, 2×2 и 2×2;
// для максимизации — 1×1, 1×1 и 3×3.
// Во втором примере для минимизации размеров,
// необходимо сделать размеры загонов 3 × 3, 3×3 и 3×3; для максимизации — 1×1, 1×1 и 7×7.
// В третьем примере для минимизации размеров,
// необходимо сделать размеры загонов 3 × 3, 3×3, 4×4 и 4×4; для максимизации — 1×1, 1×1, 1×1 и 11×11.

package main

import (
	"fmt"
)

func min(n, k int64) int64 {
	medium := n / k
	extra := n % k

	minSum := (k-extra)*medium*medium + extra*(medium+1)*(medium+1)
	return minSum
}

func max(n, k int64) int64 {
	maxSum := (k-1)*1*1 + (n-(k-1))*(n-(k-1))
	return maxSum
}

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	minSum := min(n, k)
	maxSum := max(n, k)

	fmt.Println(minSum, maxSum)

}
