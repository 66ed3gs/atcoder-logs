package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var Map map[int]map[int]int = map[int]map[int]int{}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for x := 0; x < 30; x++ {
		Map[x] = map[int]int{}
		for y := 0; y < 30; y++ {
			Map[x][y] = 0
		}
	}

	var count int
	for i := 0; i < 1000; i++ {
		for sc.Scan() {
			count += 1
			fmt.Println(count)
			in := splitInt(sc.Text())
			route := calcRoute(in[0], in[1], in[2], in[3])
			fmt.Println(strings.Join(route, ""))

			// wait for returning score
			for sc.Scan() {
				score, _ := strconv.Atoi(sc.Text())
				record(in[0], in[1], route, score)

				for x := 0; x < 30; x++ {
					for y := 0; y < 30; y++ {
						fmt.Printf("%04d,", Map[x][y])
					}
					fmt.Println()
				}
				fmt.Println()
				fmt.Println()
				break
			}
		}

	}

}

func splitInt(line string) (result []int) {
	in := strings.Split(line, " ")

	for _, v := range in {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}

	return result
}

func record(y, x int, route []string, score int) {
	avrScore := score / len(route)
	fmt.Printf("(x, y) = (%d, %d)\n", x, y)
	fmt.Printf("score: %d, average = %d\n", score, avrScore)

	for _, v := range route {
		switch v {
		case "L":
			x -= 1
			if Map[x][y] == 0 {
				Map[x][y] = avrScore
				continue
			}

			Map[x][y] = (Map[x][y] + avrScore) / 2
		case "R":
			if Map[x][y] == 0 {
				Map[x][y] = avrScore
				continue
			}

			Map[x][y] = (Map[x][y] + avrScore) / 2
			x += 1
		case "U":
			y -= 1
			if Map[x][y] == 0 {
				Map[x][y] = avrScore
				continue
			}

			Map[x][y] = (Map[x][y] + avrScore) / 2
		case "D":
			if Map[x][y] == 0 {
				Map[x][y] = avrScore
				continue
			}

			Map[x][y] = (Map[x][y] + avrScore) / 2
			y += 1
		}
	}
}

func calcRoute(sy, sx, ey, ex int) (route []string) {
	xMove := ex - sx // X Movement
	yMove := ey - sy // Y Movement
	xDist := int(math.Abs(float64(xMove)))
	yDist := int(math.Abs(float64(yMove)))

	switch {
	case xMove < 0:
		for i := 0; i < xDist; i++ {
			route = append(route, "L")
		}
	case xMove > 0:
		for i := 0; i < xDist; i++ {
			route = append(route, "R")
		}
	}

	switch {
	case yMove < 0:
		for i := 0; i < yDist; i++ {
			route = append(route, "U")
		}
	case yMove > 0:
		for i := 0; i < yDist; i++ {
			route = append(route, "D")
		}
	}

	return route
}

func calcOptimalRoute(y, x, ey, ex int, past string) (ay, ax int, action string) {

}
