package libpwgen

import (
	_ "embed"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"time"
)

//go:embed eff_large_wordlist.txt
var effWordList []byte

// populates a slice reference with a list of random eligible words.
// the total number of words generated will be numberOfWords*count.
// example 2 passwords of 3 words each will produce 6 random words from
// a total pool of eligible words from the word list.
// an eligible word meets a min length and max length requirement
func GenerateEligibleWords(words *[]string, minLength int, maxLength int, numberOfWords int, count int) (err error) {
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
		return errors.New("no eligible words found")
	}

	// Generate a random index and select a word
	for i := 0; i < numberOfWords*count; i++ {
		randomIndex := rand.Intn(len(eligible))
		*words = append(*words, eligible[randomIndex])
	}

	return nil
}

// eligibleWords will be (number of words in a password * number of passwords)
// all the passwords are stored in a single slice.
// the iterator aligns the slice with the current password being picked from the list.
// imaging two passwords of two words each
// [1,1,2,2]
// 1+i, 1+i
// 2+i, 2+i
func SelectRandomWords(eligibleWords *[]string, numberOfWords int, iterator int) (words []string, err error) {
	if len(*eligibleWords) == 0 {
		return nil, errors.New("no eligible words found")
	}

	// Generate a random index and select a word
	words = make([]string, numberOfWords)
	for i := 0; i < numberOfWords; i++ {
		rand.Seed(time.Now().UnixNano())
		words[i] = (*eligibleWords)[iterator+i]
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
