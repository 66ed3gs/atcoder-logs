package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	fmt.Println(100*calc(n)*k + 1*calc(k)*n)

}

func calc(x int) int {
	count := 0
	for i := 1; i <= x; i++ {
		count += i
	}

	return count
}
