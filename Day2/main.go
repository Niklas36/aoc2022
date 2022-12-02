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

		case "Y": // Draw
			p = p + 2
			if inputSlice[i-2] == "A" {
				p = p + 6
			} else if inputSlice[i-2] == "B" {
				p = p + 3
			} else {
				p = p + 0
			}

		case "X": // Loose
			p = p + 1
			if inputSlice[i-2] == "A" {
				p = p + 3
			} else if inputSlice[i-2] == "B" {
				p = p + 0
			} else {
				p = p + 6
			}

		case "Z": // Win
			p = p + 3
			if inputSlice[i-2] == "A" {
				p = p + 0
			} else if inputSlice[i-2] == "B" {
				p = p + 6
			} else {
				p = p + 3
			}
		}
	}
	fmt.Println(p)
}
