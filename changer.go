package main

import (
	"math/rand"
	"strings"
)

func translate(word string) string {
	stats := makeWordStats(word)

	f := func(a []int, limit int) map[int]struct{} {
		if len(a) < limit {
			limit = len(a)
		}
		m := make(map[int]struct{}, limit)
		for len(m) < limit {
			i := a[rand.Intn(len(a))]
			if _, ok := m[i]; !ok {
				m[i] = struct{}{}
			}
		}
		return m
	}

	numberPositions := f(stats.numbers, 2)
	symbolPositions := f(stats.symbols, 2)

	var out strings.Builder
	for i, c := range word {
		if _, ok := numberPositions[i]; ok {
			out.WriteRune(stats.replace[i].to)
		} else if _, ok := symbolPositions[i]; ok {
			out.WriteRune(stats.replace[i].to)
		} else {
			out.WriteRune(c)
		}
	}
	return out.String()
}
