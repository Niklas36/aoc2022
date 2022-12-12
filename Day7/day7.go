package main

import (
	_ "embed"
	"fmt"
	"path"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func main() {

	cd := ""
	fs := map[string]int{}
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "$ cd") {
			cd = path.Join(cd, strings.Fields(line)[2])
		} else if unicode.IsDigit([]rune(line)[0]) {
			var size int
			var name string
			fmt.Sscanf(line, "%d %s", &size, &name)
			fs[path.Join(cd, name)] = size
		}
	}

	fmt.Println("fs:", fs)
	ds := map[string]int{}
	for f, s := range fs {
		for d := path.Dir(f); d != "/"; d = path.Dir(d) {
			ds[d] += s
		}
		ds["/"] += s
	}

	var part1 int
	var part2 = ds["/"]
	for _, s := range ds {
		if s <= 100000 {
			part1 += s
		}
		if s+70000000-ds["/"] >= 30000000 && s < part2 {
			part2 = s
		}
	}
	fmt.Println("Part 1", part1)

	fmt.Println("Part 2", part2)
}
