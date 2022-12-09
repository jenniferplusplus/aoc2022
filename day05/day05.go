package main

import (
	"advent-of-code-2022/day05/part1"
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readInput()
	//lines := []string{
	//	"    [D]    ",
	//	"[N] [C]    ",
	//	"[Z] [M] [P]",
	//	" 1   2   3",
	//	"",
	//	"move 1 from 2 to 1",
	//	"move 3 from 1 to 3",
	//	"move 2 from 2 to 1",
	//	"move 1 from 1 to 2",
	//}

	stacks, controls, labels := part1.ParseInput(lines)
	part1.Process(stacks, controls)
	result := ""
	for i := 0; i < len(labels); i++ {
		result += part1.Top(stacks[labels[i]])
	}

	log.Println(result)
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
