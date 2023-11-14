package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var pwString string

//go:embed eff_large_wordlist.txt
var effWordList []byte

func main() {
	var minLength int
	var maxLength int
	var numberOfWords int
	var delimiter string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "minLength",
				Aliases:     []string{"min", "gt"},
				Value:       3,
				Usage:       "Minimum word length",
				Destination: &minLength,
			},
			&cli.IntFlag{
				Name:        "maxLength",
				Aliases:     []string{"max", "lt"},
				Value:       5,
				Usage:       "Maximum word length",
				Destination: &maxLength,
			},
			&cli.IntFlag{
				Name:        "words",
				Aliases:     []string{"length", "w"},
				Value:       2,
				Usage:       "Number of words to generate",
				Destination: &numberOfWords,
			},
			&cli.StringFlag{
				Name:        "delimiter",
				Aliases:     []string{"d"},
				Value:       "-",
				Usage:       "symbol to deliminate each word",
				Destination: &delimiter,
			},
		},
		Action: func(cCtx *cli.Context) error {
			words, err := GetWords(minLength, maxLength, numberOfWords)
			if err != nil {
				return cli.Exit(err, 1)
			}

			fmt.Println(words)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
