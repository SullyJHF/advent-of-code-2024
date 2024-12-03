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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(append([]int(nil), lhs...), append([]int(nil), rhs...))
	part2(append([]int(nil), lhs...), append([]int(nil), rhs...))
}

func part1(lhs, rhs []int) {
	slices.Sort(lhs)
	slices.Sort(rhs)
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
	fmt.Printf("Part 1: %d\n", total)
}

func part2(lhs, rhs []int) {
	numCounts := make(map[int]int)
	for _, num := range rhs {
		numCounts[num] += 1
	}
	similarityScore := 0
	for _, num := range lhs {
		count := numCounts[num]
		similarityScore += num * count
	}
	fmt.Printf("Part 2: %d\n", similarityScore)
}
