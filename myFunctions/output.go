package myfunctions

import (
	"os"
	"log"
)

func Output(data []string, fileName string) {

	// Open or create the file
	// os.O_RDWR means open for read/write
	// os.O_CREATE means create the file if it doesn't exist
	// os.O_TRUNC means truncate (clear) the file if it already exists
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// Write data line by line while.
	for i := range data {

		// Write line.
		_, err := file.WriteString(data[i]+"\n")
		if err != nil {
			log.Println(err)
			return
		}
	}
}
