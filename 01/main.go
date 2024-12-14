package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findRepeats(val int, list []int) int {
	count := 0
	for i := 0; i < len(list); i++ {
		if list[i] < val {
			continue
		}
		if list[i] > val {
			break
		}
		count++
	}
	return count
}

func main() {
	fmt.Println("hello world");
	dat, err := os.ReadFile("./input.txt")
	check(err)
	var listA, listB [1000]int
	// fmt.Print(string(dat))
	lines := strings.Split(string(dat), "\n")
	for i := 0; i < len(lines) - 1; i++  {
		line := lines[i];
		parts := strings.Split(line, "   ")
		a, err := strconv.Atoi(parts[0])
		check(err)
		b, err := strconv.Atoi(parts[1])
		check(err)
		listA[i] = a
		listB[i] = b
	}
	sort.Ints(listA[:])
	sort.Ints(listB[:])

	total := 0
	simscore := 0
	for i := 0; i < len(listA); i++ {
		a := listA[i]
		b := listB[i]
		diff := a - b
		if diff < 0 { diff *= -1 }
		total += diff

		repeat := findRepeats(a, listB[:])
		simscore += repeat * a
	}
	fmt.Println("total:", total)
	fmt.Println("similarity score:", simscore)
}
