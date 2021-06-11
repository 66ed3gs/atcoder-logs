package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func count(prefix int, dig int) int {
	// count := 0
	if dig == 1 {
		return 9
	}

	if prefix < dig {

	}
	return 0
}

func happy(x int64) bool {
	s := strings.Split(strconv.FormatInt(x, 10), "")
	if len(s) == 1 {
		return true
	}

	for i, v := range s[1:] {
		a, _ := strconv.Atoi(s[i])
		b, _ := strconv.Atoi(v)

		if abs(a-b) > 1 {
			return false
		}
	}

	return true
}

func solve() {
	var k int
	nextInt(&k)

	if k < 10 {
		write(k)
		return
	}

	var queue []int64
	for i := 1; i < 10; i++ {
		queue = append(queue, int64(i))
	}

	for len(queue) < k {
		write(len(queue))
		for _, v := range queue {
			if v%10 != 0 {
				queue = append(queue, 10*v+v%10-1)
			}

			queue = append(queue, 10*v+v%10)

			if v%10 != 9 {
				queue = append(queue, 10*v+v%10+1)
			}
		}
	}

	write(queue[k-1])
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

func nextStr(vars ...*string) {
	str := strings.Split(nextLine(), " ")
	for i, v := range vars {
		*v = str[i]
	}
}

func nextIntSlice(s *[]int) {
	nums := strings.Split(nextLine(), " ")

	(*s) = make([]int, len(nums))
	for i, v := range nums {
		(*s)[i], _ = strconv.Atoi(v)
	}
}

func nextStrSlice(s *[]string) {
	str := strings.Split(nextLine(), " ")

	(*s) = make([]string, len(str))
	for i, v := range str {
		(*s)[i] = v
	}
}

func sortSlice(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
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

func main() {
	solve()
	output()
}
