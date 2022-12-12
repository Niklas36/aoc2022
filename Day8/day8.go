package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	treeMap := strings.Split(input, "\n")
	treeMap1 := strings.Join(treeMap, "")
	treeMap2 := strings.Split(treeMap1, "")

	treeMap2D := make([][]string, len(treeMap))
	for i := range treeMap2D {
		treeMap2D[i] = make([]string, len(treeMap[0]))
	}

	var x int
	var y int
	counter := 0

	for j, _ := range treeMap2 {
		treeMap2D[x][y] = treeMap2[j]
		y++
		if y == len(treeMap[0]) {
			x++
			y = 0
		}
	}

	for k, _ := range treeMap2D {
		for range treeMap2D[k][0] {
			fmt.Println(k)
			counter++
			for k < len(treeMap2D[k][0])-1 {
				if treeMap2D[k+1][0] >= treeMap2D[k+2][0+k] {
					counter += 0
				} else {
					counter++
				}
				k++
			}

		}
		for range treeMap2D[0][k] {
			counter++
		}
		for range treeMap2D[len(treeMap2D)-1][k] {
			counter++
		}
		for range treeMap2D[k][len(treeMap2D)-1] {
			counter++
		}

	}
	fmt.Println(counter - 4)

}
