package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lhs := make([]int, 0)
	rhs := make([]int, 0)

	re := regexp.MustCompile(`\s+`)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := re.Split(line, 2)
		first, _ := strconv.Atoi(split[0])
		second, _ := strconv.Atoi(split[1])
		lhs = append(lhs, first)
		rhs = append(rhs, second)
	}
	slices.Sort(lhs)
	slices.Sort(rhs)
	fmt.Println(lhs)
	fmt.Println(rhs)
	total := 0
	for i := 0; i < len(lhs); i++ {
		left := lhs[i]
		right := rhs[i]
		diff := left - right
		if diff < 0 {
			diff *= -1
		}
		total += diff
	}
	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
