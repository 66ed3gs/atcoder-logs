package main

import (
  "bufio"
  "fmt"
  "os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
    sc.Scan()
    return sc.Text()
}

func main() {
  var n int
  fmt.Scanf("%d", &n)


  var users map[string]int = map[string]int{}
  for i := 1; i <= n; i++ {
    s := nextLine()

    if users[s] == 1 {
      continue
    }

    fmt.Println(i)

    users[s] = 1
  }
}


