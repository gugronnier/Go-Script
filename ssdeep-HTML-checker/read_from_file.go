package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Check if line respect the defined format
func checkFormat(txtInput string) {
	if strings.Contains(txtInput, "http://") || strings.Contains(txtInput, "https://") || strings.Contains(txtInput, "www.") {
		log.Print("[FATAL ERROR] - domains in input file don't respect defined format (see more below)")
		log.Print("\n----------------------------\n")
		usage()
	}
}


// Read input file and do request on each domain contains into it
func read_from_file(fname string) {
	// Open input file
	file, err := os.Open(fname)
	ofname := ""
	if err != nil {
		log.Fatal("failed opening file: %s\n", err)
	}

	// Read each line and them in a buffer
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	var txtline string

	for scanner.Scan() {
		txtline = scanner.Text()
		checkFormat(txtline)
		txtlines = append(txtlines, txtline)
	}

	// Close file after finished to use it
	file.Close()

	// Do request for each line in the buffer
	for _, eachline := range txtlines {
		// Create output file name forged from URL but  with raplacing "." by "_"
		ofname = "./output/" + strings.Replace(eachline, ".","_",-1) + ".html"
		_ = httpRequest(eachline, ofname)
	}
}
