package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rdr = bufio.NewReaderSize(os.Stdin, 1024*1024*1024)

func nextLine() string {
	buf := make([]byte, 0, 1024*1024*1024)
	for {
		l, p, e := rdr.ReadLine()
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

func main() {

	var n int64
	fmt.Scan(&n)

	a := strings.Split(nextLine(), " ")
	b := strings.Split(nextLine(), " ")
	c := strings.Split(nextLine(), " ")

	list := []int64{}
	for _, v := range c {
		i, _ := strconv.Atoi(v)
		b, _ := strconv.Atoi(b[i-1])
		list = append(list, int64(b))
	}

	var count int64
	for _, v1 := range list {
		for _, v2 := range a {
			i, _ := strconv.Atoi(v2)
			if v1 == int64(i) {
				count += 1
			}
		}
	}

	fmt.Println(count)

}
