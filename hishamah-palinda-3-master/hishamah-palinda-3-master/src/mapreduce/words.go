package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"strings"
	"runtime"
	"math"
	"log"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freq := make(map[string]int)
	text = strings.ToLower(text)
	words := strings.Fields(text)
	numWords := len(words)
	numCPU := runtime.NumCPU()
	skip := int(math.Ceil(float64(numWords / numCPU)));
	ch := make(chan map[string]int, numCPU)
	// 113123 words
	for i := 0; i < numCPU; i++ {
		j := (i+1) * skip
		if(i == numCPU-1){
			j = numWords
		}
		slice := words[i*skip:j]
		go func() {
			partialFreq := make(map[string]int)
			for _,word := range slice {
				lowWord := strings.TrimSuffix(word, ".")
				lowWord = strings.TrimSuffix(lowWord, ",")
				partialFreq[lowWord] += 1
			}
			ch <- partialFreq
		}()
	}

	for i:= 0; i < numCPU; i++ {	
		for k, v := range <-ch {
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
	fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
