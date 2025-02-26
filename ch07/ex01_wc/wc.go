package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"unicode"
)

var INPUT io.Reader
var FILEPATH string

func main() {
	if len(os.Args) == 1 {
		fmt.Println("interactive mode...")
		INPUT = os.Stdin
	}

	if len(os.Args) == 2 {
		path := os.Args[1]
		if f, err := os.Open(path); err == nil {
			INPUT = f
			FILEPATH = path
		} else {
			fmt.Fprintln(os.Stderr, "file does not exist")
		}
	}

	err := processInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func processInput() error {
	var words, lines, bytes, chars int
	var lastRune rune
	reader := bufio.NewReader(INPUT)
	for {
		r, s, err := reader.ReadRune()
		if errors.Is(err, io.EOF) {
			if !unicode.IsSpace(lastRune) {
				words++
			}
			break
		}
		if err != nil {
			return fmt.Errorf("reading file: %v", err)
		}
		bytes += s
		chars++

		if unicode.IsSpace(r) && !unicode.IsSpace(lastRune) {
			words++
		}

		if r == '\r' || r == '\n' {
			lines++
		}

		lastRune = r
	}

	fmt.Printf("\t%d\t%d\t%d %s\n", lines, words, chars, FILEPATH)

	return nil
}
