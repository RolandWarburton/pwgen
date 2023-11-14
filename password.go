package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

func GetWords(minLength int, maxLength int, numberOfWords int) (words []string, err error) {
	rand.Seed(time.Now().UnixNano())
	// Determine the number of CPU cores and constrain the number of workers to half of them
	numCores := runtime.NumCPU()
	numWorkers := numCores / 2

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

	if len(eligible) == 0 {
		return []string{}, errors.New("no eligible words found")
	}

	// Generate a random index and select a word
	words = make([]string, numberOfWords)
	for i := 0; i < numberOfWords; i++ {
		randomIndex := rand.Intn(len(eligible))
		randomWord := eligible[randomIndex]
		words[i] = randomWord
	}

	return words, nil
}

func getRandomSymbol() string {
	symbols := []string{"!", "@", "#", "$", "%", "^", "&"}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(symbols))
	return symbols[randomIndex]
}

func ConstructPassword(words []string, delimiter string, prepended string, appended string) string {
	password := ""
	for i, word := range words {
		password += word
		if i != len(words)-1 {
			password += delimiter
		}
	}

	// if appended is not set then randomly generate one to increase password entropy
	if appended == "" {
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(10) + 1
		randomSymbol := getRandomSymbol()
		appended = fmt.Sprintf("%s%d%s", delimiter, randomNumber, randomSymbol)
	}

	// add prepended and appended words to the password
	password = fmt.Sprintf("%s%s%s", prepended, password, appended)

	return password
}
