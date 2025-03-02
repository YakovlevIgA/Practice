package main

import (
	"fmt"
	"math"
)

func sieve(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return isPrime
}

func segmentedSieve(l, r int) []int {
	sqrtR := int(math.Sqrt(float64(r))) + 1
	isPrime := sieve(sqrtR)

	size := r - l + 1
	isCompositeRange := make([]bool, size)
	for i := 0; i < size; i++ {
		isCompositeRange[i] = false
	}
	if l == 1 {
		isCompositeRange[0] = false
	}

	for i := 2; i <= sqrtR; i++ {
		if isPrime[i] {

			start := max(i*i, l+(i-l%i)%i)

			for j := start; j <= r; j += i {
				if j != i {
					isCompositeRange[j-l] = true
				}
			}
		}
	}

	composites := []int{}
	for i := 0; i < size; i++ {
		if isCompositeRange[i] {
			if l+i != 1 {
				composites = append(composites, l+i)
			}
		}
	}

	return composites
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func countDivisors(n int) int {
	count := 0
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			if i == n/i {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func countElementsWithPrimeDivisors(arr []int) int {
	count := 0
	for _, num := range arr {
		if num <= 1 {
			continue
		}
		divisors := countDivisors(num)
		if isPrime(divisors) {
			count++
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var l, r int
	fmt.Scan(&l, &r)

	composites := segmentedSieve(l, r)
	fmt.Println(countElementsWithPrimeDivisors(composites))
}
