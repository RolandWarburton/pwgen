package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var pwString string

//go:embed eff_large_wordlist.txt
var effWordList []byte

func main() {
	rand.Seed(time.Now().UnixNano())

	minLength := 2
	maxLength := 4

	err, word := GetWord(minLength, maxLength)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(word)

	// print some data about memory used for debugging
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory usage (bytes): %d\n", memStats.TotalAlloc)

}
