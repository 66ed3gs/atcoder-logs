package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	ss := strings.Split(s, "")
	ts := strings.Split(t, "")

	if strings.Count(s, "1") != strings.Count(t, "1") {
		fmt.Println(-1)
		return
	}

	var flag []int
	for i, v := range ss {
		if v != ts[i] {
			flag = append(flag, i)
		}
	}

	var result int
	if ss[flag[0]+1] == "1" {
		switch ss[flag[0]] {
		case "0":
			p := ss[flag[0]+1:]
			fmt.Println(SliceIndex(len(p), func(i int) bool { return p[i] == "0" }))

		case "1":

		}
	} else {
		result = -1
	}

	fmt.Println(flag)
	fmt.Println(result)

}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
