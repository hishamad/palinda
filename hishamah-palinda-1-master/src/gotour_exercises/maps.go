package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordsCount := make(map[string]int)

	for _, v := range strings.Fields(s) {
		wordsCount[v] += 1
	}
	return wordsCount
}

func main() {
	wc.Test(WordCount)
}
