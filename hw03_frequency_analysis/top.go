package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func Top10(text string) []string {
	sourceSlice := strings.Fields(text)
	sourceMap := sliceToMap(sourceSlice)
	sortedPairs := sortMapByKeys(sourceMap)
	sortedSlice := sortLexicography(sortedPairs)
	return returnResult(sortedSlice)
}

func sliceToMap(sourceSlice []string) map[string]int {
	sourceMap := make(map[string]int)
	for _, val := range sourceSlice {
		sourceMap[val]++
	}
	return sourceMap
}

func sortMapByKeys(sourceMap map[string]int) PairList {
	sortedSlice := make(PairList, len(sourceMap))

	i := 0
	for key, value := range sourceMap {
		sortedSlice[i] = Pair{key, value}
		i++
	}

	sort.Sort(sort.Reverse(sortedSlice))

	return sortedSlice
}

func sortLexicography(source PairList) []string {
	result := make([]string, 0, len(source))

	for _, v := range source {
		result = append(result, v.Key)
	}

	for i := 1; i < source.Len(); i++ {
		firstRepeated, lastRepeated := i-1, i-1
		for source[i-1].Value == source[i].Value {
			lastRepeated = i
			i++
			if i == source.Len() {
				break
			}
		}

		if firstRepeated != lastRepeated {
			if lastRepeated < source.Len()-2 {
				sort.Strings(result[firstRepeated : lastRepeated+1])
			} else {
				sort.Strings(result[firstRepeated:])
			}
		}
	}

	return result
}

func returnResult(source []string) []string {
	if len(source) > 10 {
		return source[:10]
	}
	return source
}
