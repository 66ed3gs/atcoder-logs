package main

import "fmt"

func main() {
	var a, b, c int64
	fmt.Scanf("%d %d %d", &a, &b, &c)

	switch {
	case a == b:
		fmt.Println(c)
	case b == c:
		fmt.Println(a)

	case c == a:
		fmt.Println(b)
	default:
		fmt.Println(0)
	}

}
