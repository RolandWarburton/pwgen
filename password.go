package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
)

func GetWord(minLength int, maxLength int) (err error, word string) {
	// Determine the number of CPU cores and constrain the number of workers to half of them
	numCores := runtime.NumCPU()
	numWorkers := numCores / 2
	fmt.Printf("running with %d cores\n", numCores)

	// create lines from the input string
	lines := strings.Split(string(effWordList), "\n")
	linesPerWorker := len(lines) / numWorkers

	// Create a channel to collect eligible words
	eligibleWords := make(chan string, len(lines))

	// Split lines into equal slices and drop any remainder
	slices := make([][]string, numWorkers)
	for i := 0; i < numWorkers; i++ {
		start := i * linesPerWorker
		end := (i + 1) * linesPerWorker
		if end > len(lines) {
			break
		}
		slices[i] = lines[start:end]
	}

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	for i, slice := range slices {
		wg.Add(1)
		go func(_ int, slice []string) {
			defer wg.Done()
			for _, line := range slice {
				length := len(line)
				if length > minLength && length <= maxLength {
					eligibleWords <- line
				}
			}
		}(i, slice)
	}

	go func() {
		wg.Wait()
		close(eligibleWords)
	}()

	// Collect eligible words
	eligible := make([]string, 0)
	for word := range eligibleWords {
		eligible = append(eligible, word)
	}

	// Generate a random index and select a word
	randomIndex := rand.Intn(len(eligible))
	randomWord := eligible[randomIndex]

	if len(eligible) == 0 {
		return errors.New("No eligible words found."), ""
	}

	// fmt.Println("Randomly selected word:", randomWord)
	return nil, randomWord
}
