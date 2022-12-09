package part1

import (
	"log"
	"strconv"
	"strings"
)

type Span struct {
	Start int
	End   int
}

func (a Span) Contains(b Span) bool {
	if a.Start <= b.Start && a.End >= b.End {
		return true
	}
	return false
}

func ParseInput(line string) (Span, Span) {
	tokens := strings.Split(line, ",")
	a := newSpan(tokens[0])
	b := newSpan(tokens[1])

	return a, b
}

func newSpan(token string) Span {
	parts := strings.Split(token, "-")
	begin, bErr := strconv.ParseInt(parts[0], 10, 32)
	end, eErr := strconv.ParseInt(parts[1], 10, 32)
	if bErr != nil || eErr != nil {
		log.Panicf("Can't parse line (%v) (%v)", bErr, eErr)
	}

	return Span{Start: int(begin), End: int(end)}
}
