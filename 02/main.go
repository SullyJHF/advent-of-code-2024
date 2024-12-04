package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reports := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		report := make([]int, 0)
		for _, v := range split {
			num, _ := strconv.Atoi(v)
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1(&reports)
	part2(&reports)
}

func isSafe(report *[]int) bool {
	safe := true
	var lastDiff int
	for i := 0; i < len(*report)-1; i++ {
		cur := (*report)[i]
		next := (*report)[i+1]
		diff := next - cur

		if i == 0 {
			lastDiff = diff
		} else {
			if lastDiff < 0 && diff > 0 {
				return false
			}
			if lastDiff > 0 && diff < 0 {
				return false
			}
			lastDiff = diff
		}

		var absDiff int
		if diff < 0 {
			absDiff = diff * -1
		} else {
			absDiff = diff
		}

		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return safe
}

func newIsSafe(report *[]int) bool {
	for i := 0; i < len(*report); i++ {
		newReport := make([]int, 0)
		newReport = append(newReport, (*report)[:i]...)
		newReport = append(newReport, (*report)[i+1:]...)
		safe := isSafe(&newReport)
		if safe {
			return true
		}
	}
	return false
}

func part1(reports *[][]int) {
	totalSafe := 0
	for _, report := range *reports {
		safe := isSafe(&report)
		if safe {
			totalSafe++
		}
	}
	fmt.Printf("Part 1: %d\n", totalSafe)
}

func part2(reports *[][]int) {
	totalSafe := 0
	for _, report := range *reports {
		if isSafe(&report) {
			totalSafe++
		} else if newIsSafe(&report) {
			totalSafe++
		}
	}
	fmt.Printf("Part 1: %d\n", totalSafe)
}
