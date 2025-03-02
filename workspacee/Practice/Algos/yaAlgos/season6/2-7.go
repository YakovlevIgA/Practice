// https://contest.yandex.ru/contest/66793/problems/G/
package main

import "fmt"

func find(n, c int, strArr []string) int {
	lenStr := 2
	var ans int
	var checkC int
	l := 0        // left
	r := 1        // right
	for l < n-1 { // 1
		if strArr[l] == "a" { // 2
			//outerLoop:
			for r < n { // 3
				if strArr[r] == "b" { // 4
					checkC = r - l - 1 // сколько элементов между l и r
					switch {           // открыть свич
					case checkC < c:
						fmt.Println("case <")
						for checkC <= c && checkC < n-2 {
							checkC += 1
						}
						lenStr = checkC + 2
						if ans <= lenStr {
							ans = lenStr
							//break outerLoop
						}
					case checkC == c:
						fmt.Println("case =")
						lenStr = c + 2
						if ans <= lenStr {
							ans = lenStr
							//break outerLoop
						}
						fmt.Println(ans)
					case checkC > c:
						fmt.Println("case >")
						//break outerLoop
					} // закрыть свич
					l += 1
					r += 1
				} else { // закрыть if == b открыть else
					r += 1
					fmt.Printf("r=%v\n", r)
				} // 4
			} // 3
		} else { // 2
			l += 1
			fmt.Printf("l=%v\n", l)
			fmt.Println(ans)
		} // 2
	} // 1
	return ans
}

func main() {
	n := 3
	c := 1
	strArray := []string{"a", "a", "b"}
	fmt.Println(find(n, c, strArray))
}
