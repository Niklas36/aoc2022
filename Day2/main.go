package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputSlice := strings.Split(input, "")

	var p int // points
	for i, v := range inputSlice {
		switch v {

		case "A": // Rock
			if inputSlice[i+2] == "Y" { // Draw, Rock, 1
				p = p + 3 + 1
			} else if inputSlice[i+2] == "X" { // Loose, Scissors, 3
				p = p + 3
			} else { // Win, Paper, 2
				p = p + 6 + 2
			}

		case "B": // Paper
			if inputSlice[i+2] == "Y" { // Draw, Paper, 2
				p = p + 3 + 2
			} else if inputSlice[i+2] == "X" { // Loose, Rock, 1
				p = p + 1
			} else { // Win, Scissors, 3
				p = p + 6 + 3
			}

		case "C": // Scissors
			if inputSlice[i+2] == "Y" { // Draw, Scissors, 3
				p = p + 3 + 3
			} else if inputSlice[i+2] == "X" { // Loose, Paper, 2
				p = p + 2
			} else { // Win, Rock, 1
				p = p + 6 + 1
			}
		}
	}
	fmt.Println(p)
}
