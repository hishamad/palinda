package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freq := make(map[string]int)
	words := strings.Fields(text)
	ch := make(chan map[string]int, 12)
	// 113123 words
	var wg sync.WaitGroup
    for i, j := 0, 10000; i < len(words); i, j = j, j+10000 {
		if j > len(words) {
            j = len(words)
		}
		wg.Add(1)
        go func(i, j int) {
			partialFreq := make(map[string]int)
            for _,word := range words[i:j] {
				lowWord := strings.ToLower(word)
				lowWord = strings.TrimSuffix(lowWord, ".")
				lowWord = strings.TrimSuffix(lowWord, ",")
				partialFreq[lowWord] += 1
			}
			ch <- partialFreq
			defer wg.Done()
		}(i, j)
	}
	wg.Wait()
	close(ch)
	for c := range ch {
		for k, v := range c {
			freq[k] += v
		}
	}
	
	
	return freq
}
		
	


// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	content, err := ioutil.ReadFile(DataFile)
	if err != nil {
		log.Fatal(err)
	}
	data := string(content)

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
