package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var n int64
	fmt.Scan(&n)

	base2 := strconv.FormatInt(n, 2)
	slice := strings.Split(base2, "")
	dig := len(slice) - 1

	var result [][]int
	for ind := 0; ind < dig; ind++ {
		var num []string = slice

		if num[ind] == "0" {
			continue
		}

		num[ind] = "0"
		c, _ := strconv.ParseInt(strings.Join(num, ""), 2, 64)
		b := dig - ind
		a := int(n) / int(math.Pow(2, float64(b)))

		result = append(result, []int{a, b, int(c)})
	}

	var min int = 1
	for _, v := range result {
		sum := v[0] + v[1] + v[2]
		if min > sum || min == 1 {
			min = sum
		}
	}

	fmt.Println(min)
}
