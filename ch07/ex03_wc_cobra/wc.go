package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"os"
	"unicode"
)

var wTotal, lTotal, mTotal, cTotal int

// func init() {
// 	pflag.BoolP("words", "w", false, "word count")
// 	pflag.BoolP("lines", "l", false, "line count")
// 	pflag.BoolP("chars", "m", false, "char count")
// 	pflag.BoolP("bytes", "c", false, "bytes count")
// }

func start() {
	pflag.Parse()
	args := pflag.Args()

	wasSet := false
	errorExit := false

	pflag.Visit(func(pf *pflag.Flag) {
		wasSet = true
	})

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		fmt.Fprintf(os.Stderr, "binding flags: %v\n", err)
		os.Exit(1)
	}

	if !wasSet {
		viper.Set("lines", true)
		viper.Set("words", true)
		viper.Set("chars", true)
	}

	if len(args) == 0 {
		err = processInput(os.Stdin, "")
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading stdin: %v\n", err)
			errorExit = true
		}
	} else {
		for _, path := range args {
			err := ProcessFile(path)
			if err != nil {
				fmt.Println(err)
				errorExit = true
				continue
			}
		}
	}

	if len(args) > 1 {
		printRes(lTotal, wTotal, cTotal, mTotal, "total")
	}

	if errorExit {
		os.Exit(1)
	}
}

func ProcessFile(path string) error {
	if f, err := os.Open(path); err == nil {

		err = processInput(f, path)
		if err != nil {
			return fmt.Errorf("%s: processing file: %v", path, err)
		}

	} else {
		return fmt.Errorf("%s: no such file or directory", path)
	}

	return nil
}

func processInput(input io.Reader, filepath string) error {
	var words, lines, bytes, chars int
	var lastRune rune
	reader := bufio.NewReader(input)
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

	wTotal += words
	lTotal += lines
	cTotal += bytes
	mTotal += chars

	printRes(lines, words, chars, bytes, filepath)

	return nil
}

func printRes(lines, words, chars, bytes int, filepath string) {
	res := ""

	if viper.GetBool("lines") {
		res += fmt.Sprintf("\t%d", lines)
	}
	if viper.GetBool("words") {
		res += fmt.Sprintf("\t%d", words)
	}
	if viper.GetBool("chars") && !viper.GetBool("bytes") {
		res += fmt.Sprintf("\t%d", chars)
	}
	if viper.GetBool("bytes") {
		res += fmt.Sprintf("\t%d", bytes)
	}
	fmt.Println(res, " ", filepath)
}
