package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputSlice := strings.Split(input, "\n")

	//row := strings.Split(inputSlice[0], "")

	var equals []string
	var priority int
	for i := 0; i < len(inputSlice); i++ {
		rows := strings.Split(inputSlice[i], "")
		halfLength := len(rows) / 2
		firstHalf := rows[:halfLength]
		secondHalf := rows[halfLength:]

		for _, v := range firstHalf {
			for i := 0; i < len(firstHalf); i++ {
				if v == secondHalf[i] {
					equals = nil
					equals = append(equals, v)
				}
			}
		}
		value := equals[0]
		fmt.Println(value)
		valueByte := []byte(value)
		priority = priority + charToInt(valueByte[0])

		fmt.Println(priority)
	}

	/*
		halfLength := len(row) / 2
		firstHalf := row[:halfLength]
		secondHalf := row[halfLength:]
		fmt.Println(firstHalf)
		fmt.Println(secondHalf)

		for _, v := range firstHalf {
			for i := 0; i < len(firstHalf); i++ {
				if v == secondHalf[i] {
					equals = append(equals, v)
				}
			}
		}


	*/

}

func charToInt(b byte) int {
	if b >= 'a' {
		return int(b - 'a' + 1)
	} else {
		return int(b - 'A' + 27)
	}

}
