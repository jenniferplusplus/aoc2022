package part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Backpack struct {
	items []int
}

func (b *Backpack) Sum() int {
	result := 0
	for _, item := range b.items {
		result += item
	}

	return result
}

func Parse(list *[]string) []Backpack {
	result := make([]Backpack, 1)
	pack := &result[0]

	for i, s := range *list {
		if s == "" {
			var newPack Backpack
			result = append(result, newPack)
			pack = &result[len(result)-1]
			continue
		}

		v, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			log.Panicf("Couldn't parse line %v (%v)", i, err)
		}
		pack.items = append(pack.items, int(v))
	}
	return result
}

func Read(path string) []string {
	file, err := os.Open(path)
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
