package main

import (
	//"advent-of-code-2022/day02/part1"
	"advent-of-code-2022/day02/part2"
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readInput()
	//lines := []string{"A Y", "B X", "C Z"}

	var sum int
	for _, line := range lines {
		//sum += part1.Score(line)
		sum += part2.Score(line)
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
