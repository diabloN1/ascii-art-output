package myfunctions

import (
	"fmt"
	"os"
	"strings"
)

func GetArgs() (string, string, string) {
	if len(os.Args) > 4 || len(os.Args) == 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	file := ""
	args := os.Args[1:]
	str := ""
	banner := ""
	isOutput := false
	if len(args) >= 2 && len(args[0]) >= 13 && args[0][:9] == "--output=" && args[0][len(os.Args[1])-4:] == ".txt" {
		file = strings.Split(args[0], "--output=")[1]
		str = args[1]
		isOutput = true
		if len(args) == 3 {
			banner = args[2] + ".txt"
		} else {
			banner = "standard.txt"
		}
	} else if len(args) == 2 {
		str, banner = args[0], args[1]+".txt"
	} else {
		str, banner = args[0], "standard.txt"
	}

	if !isOutput && len(os.Args) == 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	return file, str, banner
}
