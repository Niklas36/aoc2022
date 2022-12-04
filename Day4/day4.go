package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {

	inputSplit := strings.Split(input, "\n")
	inputSplit1 := strings.Join(inputSplit, ",")
	inputSplit2 := strings.Split(inputSplit1, ",")
	inputSplit3 := strings.Join(inputSplit2, " ")
	inputSplit4 := strings.Split(inputSplit3, "-")
	inputSplit5 := strings.Join(inputSplit4, " ")
	inputSplit6 := strings.Split(inputSplit5, " ")

	var counter int = 0

	nums := make([]int, len(inputSplit6))

	var pairCounter int

	fmt.Println(inputSplit6)
	fmt.Println(len(inputSplit6))

	for i, _ := range inputSplit6 {
		nums[i], _ = strconv.Atoi(inputSplit6[i])
	}

	for x := 0; x < len(inputSplit6)/4; x++ {
		if nums[counter] < nums[counter+2] && nums[counter+1] >= nums[counter+2] {
			pairCounter++
		} else if nums[counter] == nums[counter+2] {
			pairCounter++
		} else if nums[counter+2] < nums[counter] && nums[counter+3] >= nums[counter] {
			pairCounter++
		}

		counter += 4

	}

	fmt.Println(pairCounter)

}
