package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// If there is an error, exit of the programe and print the error
func check(e error) {
	if e != nil {
                log.Fatal(e)
        }
}

// If there is an error, print the error but not leave the program
func warning(e error) int {
	if e != nil {
//		log.Print(e)
		return 1
	}
	return 0
}

// Make HTTP request for specified URL and put output in file
func httpRequest(uri string, outputfname string) bool {
        url := "http://" + uri + "/"
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create HTTP request and modify User-Agent before sending
	request, err := http.NewRequest("GET", url, nil)
	check(err)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:64.0) Gecko/20100101 Firefox/64.0")

	// Make HTTP Get request
        response, err := client.Do(request)
	url_err := warning(err)
	if url_err == 1 {
		return false
	}
	defer response.Body.Close()

	// Create output file
	outFile, err := os.Create(outputfname)
        check(err)
        defer outFile.Close()

        // Copy data from HTTP response to output file
        _, err = io.Copy(outFile, response.Body)
	check(err)

	// Close opened file before exit
	outFile.Close()

	return true
}
