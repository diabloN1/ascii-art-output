package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"asciiArtOutput/myFunctions"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}
	file := ""
	if len(os.Args[1]) < 13 || os.Args[1][:9] != "--output=" || os.Args[1][len(os.Args[1])-4:] != ".txt" {
		fmt.Println("Invalid --output Flag")
		return
	} else {
		file = strings.Split(os.Args[1], "--output=")[1]
	}
	banner := os.Args[3]+".txt"
	standard, err := myfunctions.Read(banner)
	if err != nil {
		return
	}
	asciiChars := myfunctions.BytesToAsciiMap(standard)
	result, err := myfunctions.WriteResult(asciiChars)
	if err != nil {
		log.Println(err)
		return
	}
	myfunctions.Output(result, file)
}
