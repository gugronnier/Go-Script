package main

import (
	"os"
	"fmt"
	"strings"
)

// Usage function, describe how to use the program
func usage() {
    fmt.Println("Usage:   ./cybersquatting-webscraper input_file.txt client_domain.com \n")
	fmt.Println("	input_file.txt       contains domains must be tested, one per line")
	fmt.Println("	   all domains must use the same format without \"http(s)://www.\" ")
	fmt.Println("	client_domain.com    mustn't include \"http(s)://www.\" ")
	os.Exit(0)
}

func main() {
        // File contains targeted URL must be enter as argument when calling the program
        args := os.Args[1:]
        if len(args) != 2 {
                usage()
        }
        inputfname := args[0]
	refDomain := args[1]
	if strings.Contains(refDomain, "http://") || strings.Contains(refDomain, "www.") || strings.Contains(refDomain, "https://") {
		usage()
	}

	// Create directory for HTML output
	_ = os.Mkdir("output", 0777)

	// Setup Reference file
	refDomainFile := "./output/" + strings.Replace(refDomain, ".","_",-1) + ".html"
	if !httpRequest(refDomain, refDomainFile) {
		fmt.Println("[FATAL ERROR] - Fail to create reference file, check client domain entered before retry")
	}

	// Read file line by line and get the page for each URL
	read_from_file(inputfname)

	// use ssdeep custom algorithm to detect which are the same and which are different
	compareResult(refDomainFile)
}
