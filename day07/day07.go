package main

import (
	"advent-of-code-2022/day07/part1"
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readInput()
	//lines := []string{
	//	"$ cd /",
	//	"$ ls",
	//	"dir a",
	//	"14848514 b.txt",
	//	"8504156 c.dat",
	//	"dir d",
	//	"$ cd a",
	//	"$ ls",
	//	"dir e",
	//	"29116 f",
	//	"2557 g",
	//	"62596 h.lst",
	//	"$ cd e",
	//	"$ ls",
	//	"584 i",
	//	"$ cd ..",
	//	"$ cd ..",
	//	"$ cd d",
	//	"$ ls",
	//	"4060174 j",
	//	"8033020 d.log",
	//	"5626152 d.ext",
	//	"7214296 k",
	//}

	root := part1.Parse(lines)

	//go cwd.Size(100000, match)
	//
	//for size := range match {
	//	sum += size
	//}

	filter := func(values chan int64, max int64) chan int64 {
		out := make(chan int64)
		go func() {
			defer close(out)
			for v := range values {
				if v <= max {
					out <- v
				}
			}
		}()
		return out
	}

	reduce := func(values chan int64) chan int64 {
		out := make(chan int64)
		go func() {
			defer close(out)
			sum := int64(0)
			for v := range values {
				sum += v
			}
			out <- sum
		}()
		return out
	}

	values := make(chan int64)
	filtered := filter(values, 100_000)
	filteredSum := reduce(filtered)
	sum := root.Size(values)
	log.Println("Total:", sum)
	log.Println("Filtered:", <-filteredSum)
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
