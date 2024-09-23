package myfunctions

import (
	"fmt"
	"os"
	"strings"
)

func GetArgs() (string, string, string) {
	file := ""
	str := ""
	banner := ""

	// Checking len qrgs validity.
	if len(os.Args) > 4 || len(os.Args) == 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	args := os.Args[1:]
	isOutput := false

	if len(args) >= 2 && len(args[0]) >= 13 && args[0][:9] == "--output=" && args[0][len(os.Args[1])-4:] == ".txt" {

		// Case flag to output exists
		file = strings.Split(args[0], "--output=")[1]
		str = args[1]
		isOutput = true

		// Handling Flag if exists
		if len(args) == 3 {
			banner = args[2] + ".txt"
		} else {
			banner = "standard.txt"
		}
	} else if len(args) == 2 {
		// Case no output flag.
		str, banner = args[0], args[1]+".txt"
	} else {
		// Case no banner no output flag.
		str, banner = args[0], "standard.txt"
	}

	// Case 3 args and could not find flag.
	if !isOutput && len(args) == 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	return file, str, banner
}
