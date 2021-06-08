package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve() {

}

var (
	reader = bufio.NewReaderSize(os.Stdin, 1e8)
	writer = bufio.NewWriter(os.Stdout)
)

func nextLine() string {
	buf := make([]byte, 0)
	for {
		l, p, e := reader.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func nextInt(vars ...*int) {
	nums := strings.Split(nextLine(), " ")
	for i, v := range vars {
		*v, _ = strconv.Atoi(nums[i])
	}
}

func nextIntSlice(s *[]int) {
	nums := strings.Split(nextLine(), " ")

	(*s) = make([]int, len(nums))
	for i, v := range nums {
		(*s)[i], _ = strconv.Atoi(v)
	}
}

func write(s ...interface{}) {
	fmt.Fprintln(writer, s...)
}

func writeIntSlice(s []int) {
	output := make([]interface{}, len(s))
	for i, v := range s {
		output[i] = v
	}
	fmt.Fprintln(writer, output...)
}

func output() {
	_ = writer.Flush()
}

func sortSlice(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func max(A ...int) int {
	n := A[0]
	for _, v := range A {
		if n < v {
			n = v
		}
	}
	return n
}

func min(A ...int) int {
	n := A[0]
	for _, v := range A {
		if n > v {
			n = v
		}
	}
	return n
}

func sum(A ...int) int {
	sum := 0
	for _, v := range A {
		sum += v
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func main() {
	solve()
	output()
}
