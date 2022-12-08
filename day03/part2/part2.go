package part2

import (
	"advent-of-code-2022/day03/part1"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
)

func Score(set mapset.Set[rune]) int {
	if set.Cardinality() > 1 {
		log.Panicf("too many dupes (%v)", set.String())
	}

	r, ok := set.Pop()
	if !ok {
		log.Panicf("no dupes")
	}

	return part1.Score(r)
}

func Intersect(lines []string) mapset.Set[rune] {
	intersection := toSet(lines[0])
	for i := 1; i < len(lines); i++ {
		lineSet := toSet(lines[i])
		intersection = intersection.Intersect(lineSet)
	}

	return intersection
}

func toSet(line string) mapset.Set[rune] {
	lineSet := mapset.NewSet[rune]()
	for _, c := range line {
		lineSet.Add(c)
	}

	return lineSet
}
