package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	mulMatches := make([][]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		result := re.FindAllStringSubmatch(line, -1)
		mulMatches = append(mulMatches, result...)
		// fmt.Println(result)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(&mulMatches)
}

func part1(mulMatches *[][]string) {
	total := 0
	for _, v := range *mulMatches {
		x, _ := strconv.Atoi(v[1])
		y, _ := strconv.Atoi(v[2])

		total += x * y
	}
	fmt.Printf("Part 1: %d\n", total)
}

func part2() {

}
