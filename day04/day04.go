package main

import (
	"advent-of-code-2022/day04/part1"
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readInput()
	//lines := []string{
	//	"2-4,6-8",
	//	"2-3,4-5",
	//	"5-7,7-9",
	//	"2-8,3-7",
	//	"6-6,4-6",
	//	"2-6,4-8",
	//}

	overlaps := 0
	for _, line := range lines {
		a, b := part1.ParseInput(line)
		if a.Overlaps(b) {
			overlaps++
		}
	}

	log.Println(overlaps)
}

func readInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Panicf("Couldn't open input file (%v)", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
