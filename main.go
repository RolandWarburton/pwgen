package main

import (
	"syscall/js"

	"github.com/rolandwarburton/pwgen/libpwgen"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("pwgen", js.FuncOf(pwgen))
	<-c
}

func pwgen(this js.Value, p []js.Value) interface{} {
	var wordsList []string
	minLength := 3
	maxLength := 5
	numberOfWords := 2
	count := 1
	delimiter := "-"
	prepend := ""
	appended := ""

	// generate the words list
	err := libpwgen.GenerateEligibleWordsWASM(&wordsList, minLength, maxLength, numberOfWords, count)
	if err != nil {
		return js.ValueOf("")
	}
	words, err := libpwgen.SelectRandomWords(&wordsList, numberOfWords, 0*numberOfWords)
	if err != nil {
		return js.ValueOf("")
	}
	pw := libpwgen.ConstructPassword(words, delimiter, prepend, appended)
	return js.ValueOf(pw)
}
