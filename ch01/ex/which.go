package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No arguments")
		os.Exit(1)
	}
	file := os.Args[1]

	envPath := os.Getenv("PATH")
	dirs := filepath.SplitList(envPath)
	for _, dir := range dirs {
		fullPath := filepath.Join(dir, file)
		info, err := os.Stat(fullPath)
		if err != nil {
			// does not exist
			continue
		}

		mode := info.Mode()
		if !mode.IsRegular() {
			// not a file
			continue
		}

		if mode&0111 != 0 {
			fmt.Println(fullPath)
		}
	}
}
