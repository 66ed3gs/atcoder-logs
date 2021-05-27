package main

import (
	"fmt"
)

func main() {
	var a, b, k int64
	fmt.Scanf("%d %d %d", &a, &b, &k)

	calc(a, b, k)
}

func calc(a int64, b int64, k int64) {
	var i int64
	for i = 1; i == i; i++ {
		if b == 0 {
			for i = 1; i <= a; i++ {
				fmt.Print("a")
			}

			for i = 1; i <= a; i++ {
				fmt.Print("a")
			}

			return
		}

		dec := a + b - int64(i)
		if k-dec <= 0 {
			fmt.Print("a")
			a -= 1
			b -= i
			calc(a, b, k)
			return
		}
		k -= dec
		fmt.Print("b")
	}
}
