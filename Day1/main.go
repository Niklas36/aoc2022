package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed calories.txt
var input string

func main() {
	err := readCalories()
	if err != nil {
		fmt.Errorf("cannot read calories: %w", err)
	}
}

func readCalories() error {
	caloriesListString := strings.Split(input, "\n")
	var tmp []string
	for _, calorie := range caloriesListString {
		tmp = append(tmp, calorie)
	}

	caloriesListInt := make([]int, len(caloriesListString))

	for i, s := range caloriesListString {
		caloriesListInt[i], _ = strconv.Atoi(s)
	}

	sum := 0
	var caloriesSums []int
	for _, v := range caloriesListInt {
		if v != 0 {
			sum += v
			caloriesSums = append(caloriesSums, sum)
		} else {
			sum = 0
		}
	}
	sort.Slice(caloriesSums, func(i, j int) bool {
		return caloriesSums[i] < caloriesSums[j]
	})

	highestSum := caloriesSums[len(caloriesSums)-1]
	fmt.Println("Highest calories: ", highestSum)

	sumTop3 := caloriesSums[len(caloriesSums)-1] + caloriesSums[len(caloriesSums)-2] + caloriesSums[len(caloriesSums)-3]
	fmt.Println("Sum Top 3 calories: ", sumTop3)
	return nil
}
