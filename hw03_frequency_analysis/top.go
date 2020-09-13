package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"sort"
	"strings"
)

type wordCounter struct {
	name  string
	value int
}

func getWordCountersDict(arr []string) map[string]int {
	dict := map[string]int{}

	for _, v := range arr {
		dict[v]++
	}

	return dict
}

func getSortedWordCounters(dict map[string]int) []wordCounter {
	counters := make([]wordCounter, 0, len(dict))

	for k, v := range dict {
		counters = append(counters, wordCounter{name: k, value: v})
	}

	sort.Slice(counters, func(i, j int) bool {
		return counters[i].value > counters[j].value
	})

	return counters
}

func Top10(str string) []string {
	arr := strings.Fields(str)
	dict := getWordCountersDict(arr)
	counters := getSortedWordCounters(dict)
	res := make([]string, 0, len(counters))

	for _, v := range counters {
		res = append(res, v.name)
		if len(res) == 10 {
			break
		}
	}

	return res
}
