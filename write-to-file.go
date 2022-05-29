package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteToFile() {
	filePath := "/Users/garrettsweeney/git/garrett-learns-go/write-to-file.txt"
	f, err := os.Create(filePath)
	check(err)

	defer f.Close()

	bytesWritten, err := f.WriteString("I can write to a file\n")
	check(err)
	fmt.Printf("Wrote %d bytes to %s\n", bytesWritten, filePath)

	f.Sync()
}

// func main() {
//     WriteToFile()
// }
