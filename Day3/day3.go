package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	inputSlice := strings.Split(input, "\n")

	var equals []string
	//var priority int

	//	for i := 0; i < 3; i++ {
	fmt.Println(inputSlice[0])
	fmt.Println(inputSlice[1])
	fmt.Println(inputSlice[2])

	sort.SliceStable(inputSlice[0:2], func(i, j int) bool {
		return len(inputSlice[i]) > len(inputSlice[j])
	})

	firstRow := inputSlice[0]
	secondRow := inputSlice[1]
	thirdRow := inputSlice[2]
	firstRowSlice := strings.Split(firstRow, "")
	secondRowSlice := strings.Split(secondRow, "")
	thirdRowSlice := strings.Split(thirdRow, "")

	fmt.Println(inputSlice[0])
	fmt.Println(inputSlice[1])
	fmt.Println(inputSlice[2])

	for _, v := range thirdRowSlice {
		for i := 0; i < len(thirdRowSlice); i++ {
			if v == firstRowSlice[i] || v == secondRowSlice[i] {
				equals = append(equals, v)

			}
		}
		fmt.Println(equals)
	}

	/*
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


	*/

}

/*
   func charToInt(b byte) int {
   	if b >= 'a' {
   		return int(b - 'a' + 1)
   	} else {
   		return int(b - 'A' + 27)
   	}
   }

*/
