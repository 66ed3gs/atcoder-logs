package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	slice := strings.Split(s, "")

	var a, b []int
	for i, v := range slice {
		switch v {
		case "o":
			a = append(a, i)
			b = append(b, i)
		case "?":
			b = append(b, i)
		}
	}

	if len(a) == 4 {
		fmt.Println(1)
	} else if len(a) > 4 {
		fmt.Println(0)
	} else {

		var list []string
		for _, v := range b {
			for i := 0; i < len(a); i++ {
				nums := a
				nums = append(nums[:i+1], nums[i:]...)
				nums[i] = v

				list = append(list, strings.Trim(strings.Replace(fmt.Sprint(nums), " ", "", -1), "[]"))
			}
		}

		fmt.Println(list)
		/*

			m := make(map[string]struct{})
			newList := make([]string, 0)

			for _, element := range list {
				if _, ok := m[element]; !ok {
					m[element] = struct{}{}
					newList = append(newList, element)
				}
			}

			fmt.Println(len(newList))

		*/
	}

}
