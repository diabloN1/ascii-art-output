package main

import (
	"fmt"

	myfunctions "asciiArtOutput/myFunctions"
)

func main() {
	file, str, banner := myfunctions.GetArgs()
	standard, err := myfunctions.Read(banner)
	if err != nil {
		return
	}
	asciiChars := myfunctions.BytesToAsciiMap(standard)
	result, err := myfunctions.WriteResult(str, asciiChars)
	if err != nil {
		fmt.Println(err)
		return
	}
	if file != "" {
		myfunctions.Output(result, file)
	} else {
		myfunctions.PrintResult(result)
	}
}
