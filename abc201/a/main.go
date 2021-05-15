package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	nums := []int{a, b, c}
	sort.Sort(sort.IntSlice(nums))

	if nums[2]-nums[1] == nums[1]-nums[0] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
