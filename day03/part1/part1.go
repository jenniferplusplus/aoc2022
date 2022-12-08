package part1

import (
	mapset "github.com/deckarep/golang-set/v2"
)

func ScoreDupes(line string) int {
	itemSet := mapset.NewSet[rune]()
	i := len(line) / 2
	for _, c := range line[:i] {
		itemSet.Add(c)
	}

	for _, c := range line[i:] {
		if itemSet.Contains(c) {
			return Score(c)
		}
	}

	return 0
}

func Score(item rune) int {
	if item >= 65 && item <= 90 {
		return int(item - 38)
	}
	if item >= 97 && item <= 122 {
		return int(item - 96)
	}
	return 0
}
