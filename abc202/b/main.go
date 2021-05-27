package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	slice := strings.Split(s, "")

	for i := range slice {
		str := slice[len(slice)-i-1]

		switch str {

		case "6":
			fmt.Print(9)

		case "9":
			fmt.Print(6)

		default:
			fmt.Print(str)
		}
	}

	fmt.Println()

}
