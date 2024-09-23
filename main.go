package main

import (
	"fmt"

	myfunctions "asciiArtOutput/myFunctions"
)

func main() {
	// Get args.
	file, str, banner := myfunctions.GetArgs()

	// Read banner file.
	standard, err := myfunctions.Read(banner)
	if err != nil {
		return
	}

	// Bytes to asciiMap.
	asciiChars := myfunctions.BytesToAsciiMap(standard)
	result, err := myfunctions.WriteResult(str, asciiChars)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output data.
	if file != "" {
		myfunctions.Output(result, file)
	} else {
		myfunctions.PrintResult(result)
	}
}
