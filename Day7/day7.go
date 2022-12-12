package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	var dataSize int
	var total int
	for _, line := range strings.Split(input, "\n") {
		fmt.Sscanf(line, "%d", &dataSize)

		if dataSize <= 100000 {
			total += dataSize
			fmt.Println("Total:", dataSize)
		}
		dataSize = 0
	}
	fmt.Println("Ergebnis:", total)

}
