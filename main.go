// fastfind is a command-line tool to quickly search for files and directories
// matching a given term within a specified directory (or current directory by default).
// It highlights directories in bold and skips .git directories during the search.	
// Usage: fastfind <search_term> [path]

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
)

func main() {
	// Expecting: fastfind <term> [optional_path]
	if len(os.Args) < 2 {
		fmt.Println("Usage: fastfind <search_term> [path]")
		return
	}

	searchTerm := strings.ToLower(os.Args[1])
	
	// Default to current directory if no path is provided
	root := "."
	if len(os.Args) > 2 {
		root = os.Args[2]
	}

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		// Silently skip errors (like permission denied) to keep output clean
		if err != nil {
			return nil 
		}

		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		if strings.Contains(strings.ToLower(d.Name()), searchTerm) {
			if d.IsDir() {
				fmt.Printf("[DIR]  %s%s%s\n", Bold, path, Reset)
			} else {
				fmt.Printf("[FILE] %s\n", path)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}