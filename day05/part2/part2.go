package part2

import (
	"advent-of-code-2022/day05/part1"
	"github.com/oleiade/lane/v2"
)

func Process(stacks map[int32]*lane.Stack[rune], controls []part1.Control) {
	for _, c := range controls {
		buffer := make([]rune, 0, c.Count)
		for i := 0; i < c.Count; i++ {
			val, popped := stacks[c.Src].Pop()
			if popped {
				buffer = append(buffer, val)
			}
		}

		for b := len(buffer) - 1; b >= 0; b-- {
			stacks[c.Dest].Push(buffer[b])
		}
	}
}
