package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b int64
	fmt.Scanf("%d %d", &a, &b)

	gcd := gcd(a, b)
	a2 := big.NewInt(a / gcd)
	b2 := big.NewInt(b / gcd)

	num := big.NewInt(gcd)
	num.Mul(num, a2)
	num.Mul(num, b2)

	if num.Cmp(big.NewInt(1000000000000000000)) == 1 {
		fmt.Println("Large")
		return
	}

	fmt.Println(num)
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
