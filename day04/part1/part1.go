package part1

import (
	"github.com/oleiade/lane/v2"
	"log"
)

type offset struct {
	col int
	id  rune
}

func ParseInput(lines []string) {
	init, _ := splitInput(lines)
	stacks := parseInit(init)
	head, present := stacks['5'].Head()
	log.Print(string(head), present)
}

func splitInput(lines []string) (init []string, control []string) {
	i := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
	}

	return lines[:i], lines[i+1:]
}

func parseInit(lines []string) map[int32]*lane.Stack[rune] {
	labels := lines[len(lines)-1]
	values := lines[:len(lines)-1]
	stacks := map[int32]*lane.Stack[rune]{}
	offsets := make([]offset, 0, 9)

	// Initialize empty stacks
	for col, id := range labels {
		if id != ' ' {
			offsets = append(offsets, offset{col, id})
			stacks[id] = lane.NewStack[rune]()
		}
	}

	// Fill the stacks
	for i := len(values) - 1; i >= 0; i-- {
		line := values[i]
		for _, o := range offsets {
			if line[o.col] != ' ' {
				stacks[o.id].Push(rune(line[o.col]))
			}
		}
	}

	return stacks
}
