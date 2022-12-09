package main

import (
	"advent-of-code-2022/day06/part1"
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readInput()
	//lines := []string{
	//	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	//}

	log.Println(part1.LocateMarkerStart(lines[0], 14))
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
