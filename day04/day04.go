package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readInput()
	//lines := []string{

	//}

	log.Println(lines[:10])
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
