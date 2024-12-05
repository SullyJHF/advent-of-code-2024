package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	allLines := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		allLines = append(allLines, line)
		result := re.FindAllStringSubmatch(line, -1)
		mulMatches = append(mulMatches, result...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(&mulMatches)

	content := strings.Join(allLines, " ")
	chars := make([]string, 0)
	allChars := make([]string, 0)
	enabled := true
	for i := 0; i < len(content); i++ {
		curChar := content[i]
		allChars = append(allChars, string(curChar))
		if len(allChars) > 7 {
			if strings.Join(allChars[i-7:i], "") == "don't()" && enabled {
				enabled = false
				chars = chars[:len(chars)-7]
			}
		}
		if len(allChars) > 4 {
			if strings.Join(allChars[i-4:i], "") == "do()" && !enabled {
				enabled = true
			}
		}
		if enabled {
			chars = append(chars, string(curChar))
		}

	}

	mulMatches = re.FindAllStringSubmatch(strings.Join(chars, ""), -1)
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
