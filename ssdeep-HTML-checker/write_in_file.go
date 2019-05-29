package main

import (
	"bufio"
	"os"
)

// Take file where write as parameter and string to write
// And write in the file without overwriting what is written in before
func writeInFile(input string, ofname string) {
        outputFile, err := os.Open(ofname)
	check(err)
	outputFile.Sync()
        w := bufio.NewWriter(outputFile)
        _, err = w.WriteString(input + "\n")
        check(err)
	outputFile.Close()
        w.Flush()
}
