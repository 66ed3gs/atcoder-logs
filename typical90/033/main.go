package main

import (
  "fmt"
)

func main() {
  var h, w int
  fmt.Scanf("%d %d", &h, &w)

  if h == 1 || w == 1 {
    fmt.Println(h * w)
    return
  }

  fmt.Println(int((h+1)/2) * int((w+1)/2))
}
