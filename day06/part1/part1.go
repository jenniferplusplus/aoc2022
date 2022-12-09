package part1

import (
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"strings"
)

func LocatePacketStart(line string) int {
	i := 4
	for ; i < len(line); i++ {
		chars := strings.SplitAfter(line[i-4:i], "")
		set := mapset.NewSet[string](chars...)
		log.Println(set.String())
		if set.Cardinality() == 4 {
			return i
		}
	}

	return i
}

// LocateMarkerStart part 2, same, but longer
func LocateMarkerStart(line string, markerSize int) int {
	i := markerSize
	for ; i < len(line); i++ {
		chars := strings.SplitAfter(line[i-markerSize:i], "")
		set := mapset.NewSet[string](chars...)
		log.Println(set.String())
		if set.Cardinality() == markerSize {
			return i
		}
	}

	return i
}
