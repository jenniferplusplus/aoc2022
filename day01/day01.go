package main

import (
	"advent-of-code-2022/day01/part1"
	"advent-of-code-2022/day01/part2"
	"log"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	topN, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		log.Panic("Both args are required")
	}
	list := part1.Read(input)
	packs := part1.Parse(&list)

	top := part2.ClosedOrderedQueue[int]{}
	top.Init(topN)
	for _, pack := range packs {
		top.Push(pack.Sum())
	}

	sum := 0
	for {
		v, found := top.Pop()
		if !found {
			break
		}
		sum += v
	}

	log.Println(sum)
}
