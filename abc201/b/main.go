package main

import (
	"fmt"
	"sort"
)

func main() {
	var c int
	fmt.Scan(&c)

	var name string
	var h int
	m := make(map[string]int)
	for i := 0; i < c; i++ {
		fmt.Scanf("%s %d", &name, &h)
		m[name] = h
	}

	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(n[a[1]][0])
}
