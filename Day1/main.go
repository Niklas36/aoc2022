package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Errorf("cannot get current directory: %w", err)
	}

	fmt.Println(dir)

	readCalories(dir)

}

func readCalories(dir string) {
	fname := filepath.Join(dir, "calories.txt")

	file, err := os.Open(fname)
	if err != nil {
		fmt.Errorf("cannot open calories file: %w", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		fmt.Errorf("cannot read calories: %w", err)
	}

	calories := string(buf)

	fmt.Println(calories)

	caloriesListString := strings.Split(calories, "\n")
	var tmp []string
	for _, calorie := range caloriesListString {
		tmp = append(tmp, calorie)

	}

	caloriesListInt := make([]int, len(caloriesListString))

	for i, s := range caloriesListString {
		caloriesListInt[i], _ = strconv.Atoi(s)
	}

	fmt.Println(caloriesListInt)

	sum := 0
	highestSum := 0
	secondHighest := 0
	thirdHighest := 0

	for _, v := range caloriesListInt {
		if v != 0 {
			sum += v

		} else {
			sum = 0

		}
		if sum > highestSum {
			highestSum = sum

		} else if sum > secondHighest && sum < highestSum {
			secondHighest = highestSum

		} else if sum > thirdHighest && sum < highestSum {
			thirdHighest = sum
		}

	}

	fmt.Println(highestSum)
	fmt.Println(secondHighest)
	fmt.Println(thirdHighest)

	fmt.Println("Summe :", highestSum+secondHighest+thirdHighest)

}
