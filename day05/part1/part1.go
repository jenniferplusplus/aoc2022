package part1

import (
	"github.com/oleiade/lane/v2"
	"log"
	"strconv"
	"strings"
)

type offset struct {
	col int
	id  rune
}

type Control struct {
	Src   rune
	Dest  rune
	Count int
}

func Top(stack *lane.Stack[rune]) string {
	head, exists := stack.Head()
	if exists {
		return string(head)
	}
	return ""
}

func Process(stacks map[int32]*lane.Stack[rune], controls []Control) {
	for _, c := range controls {
		for i := 0; i < c.Count; i++ {
			val, popped := stacks[c.Src].Pop()
			if popped {
				stacks[c.Dest].Push(val)
			} else {
				log.Println(val)
			}
		}
	}
}

func ParseInput(lines []string) (map[int32]*lane.Stack[rune], []Control, []rune) {
	init, ctl := splitInput(lines)
	offsets := parseOffsets(init[len(init)-1])
	stacks := parseInit(init, offsets)
	controls := parseControl(ctl)

	log.Println(stacks)
	log.Println(controls[:4])

	labels := make([]rune, 0, len(offsets))
	for _, o := range offsets {
		labels = append(labels, o.id)
	}

	return stacks, controls, labels
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

func parseOffsets(line string) []offset {
	offsets := make([]offset, 0)
	for col, id := range line {
		if id != ' ' {
			offsets = append(offsets, offset{col, id})
		}
	}

	return offsets
}

func parseInit(lines []string, offsets []offset) map[int32]*lane.Stack[rune] {
	values := lines[:len(lines)-1]
	stacks := map[int32]*lane.Stack[rune]{}

	// Initialize empty stacks
	for _, o := range offsets {
		stacks[o.id] = lane.NewStack[rune]()
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

// move 8 from 7 to 1
// move 9 from 1 to 9
// 0	1 2    3 4  5
func parseControl(lines []string) []Control {
	controls := make([]Control, 0, len(lines))
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		count, err := strconv.ParseInt(tokens[1], 10, 32)
		if err != nil {
			log.Panicf("couldn't parse instruction line %v (%v)", i, err)
		}
		controls = append(controls, Control{Count: int(count), Src: rune(tokens[3][0]), Dest: rune(tokens[5][0])})
	}

	return controls
}
