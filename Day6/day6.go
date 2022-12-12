package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {

	marker := input[:14]

	for i := len(marker); i < len(input); i++ {
		if uniqueCharacters(marker) {
			fmt.Println("unique marker: ", i)
			break
		} else {
			marker = marker[1:] + input[i:i+1]
		}
	}

}

func uniqueCharacters(s string) bool {

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i:i+1] == s[j:j+1] {
				return false
			}
		}
	}
	return true
}
