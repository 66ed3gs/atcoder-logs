package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() (int64, int) {
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	a, _ := strconv.ParseInt(s[0], 10, 64)
	b, _ := strconv.Atoi(s[1])

	return a, b
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	var friends map[int64]int = map[int64]int{}
	for i := 0; i < n; i++ {
		a, b := nextLine()
		friends[a] += b
	}

	i := 0
	sl := make([]int64, len(friends))
	for key := range friends {
		sl[i] = key
		i += 1
	}

	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})

	money := k
	var now int64 = 0
	for _, key := range sl {
		if key-now > int64(money) {
			fmt.Println(now + int64(money))
			return
		}

		money += friends[key]
	}

	fmt.Println(now + int64(money))
}
