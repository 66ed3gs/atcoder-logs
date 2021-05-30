package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() []int {
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	var line []int = []int{}
	for i, v := range s {
		n, _ := strconv.Atoi(v)
		line[i] = n
	}

	return line
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	var area map[int][]int = map[int][]int{}
	for i := 0; i < n; i++ {
		area[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		area[i] = nextLine()
	}

	var min int = int(math.MaxInt32)
	for o := 0; o < (n - k + 1); o++ {
		for p := 0; p < (n - k + 1); p++ {
			nums := []int{}
			for q := 0; q < k; q++ {
				nums = append(nums, area[o+q][p:p+k]...)
			}

			md := calcMedian(nums)
			if min > md {
				min = md
			}
		}

	}

	fmt.Println(min)
}

func calcMedian(l []int) int {
	sort.Slice(l, func(i, j int) bool {
		return l[i] > l[j]
	})

	return l[len(l)/2]
}
