package main

import (
	"fmt"
	"sort"
	"strings"
)

type anagramGroup struct {
	first string
	words []string
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	anagrams := FindAnagrams(words)

	for key, group := range anagrams {
		fmt.Printf("%q: %q\n", key, group)
	}
}

func normalize(s string) string {
	r := []rune(strings.ToLower(s))
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func FindAnagrams(words []string) map[string][]string {
	groups := make(map[string]*anagramGroup)

	for _, word := range words {
		lower := strings.ToLower(word)
		key := normalize(lower)

		if g, exists := groups[key]; exists {
			g.words = append(g.words, lower)
		} else {
			groups[key] = &anagramGroup{
				first: lower,
				words: []string{lower},
			}
		}
	}

	// Формируем окончательный результат
	result := make(map[string][]string)
	for _, g := range groups {
		if len(g.words) >= 2 {
			sort.Strings(g.words)
			result[g.first] = g.words
		}
	}
	return result
}
