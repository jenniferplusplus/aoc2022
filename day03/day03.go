package main

import (
	//"advent-of-code-2022/day03/part1"
	"advent-of-code-2022/day03/part2"
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readInput()
	//lines := []string{
	//	"vJrwpWtwJgWrhcsFMMfFFhFp",
	//	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	//	"PmmdzqPrVvPwwTWBwg",
	//	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	//	"ttgJtRGJQctTZtZT",
	//	"CrZsJsPPZsGzwwsLwLmpwMDw",
	//}

	var sum int
	for i := 0; i < len(lines); i += 3 {
		dupes := part2.Intersect(lines[i : i+3])
		sum += part2.Score(dupes)
	}

	log.Println(sum)
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
