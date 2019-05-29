package main

import (
	"os"
	"io/ioutil"
	"bufio"
	"strings"
)

// Clean the buffer to keep only the domain
func cleanString(stin string) string {
	step1 := strings.Replace(stin, "./output/", "", -1)
	step2 := strings.Replace(step1, ".html", "", -1)
	step3 := strings.Replace(step2, "_", ".", -1)
	return step3
}

// Check all files in a directory in order to sort which matching with reference and which not
func compareResult(clientRefDomainFile string){
	files, err := ioutil.ReadDir("./output")
	check(err)

	// Create final output directory that will contains 2 files
	// 	client_owner_domains.txt  -  that contains only domains match
	// 				     with the client reference domain
	// 	suspicious_domains.txt    -  that contains domains doesn't match
	//				     with the client reference domain
	os.Mkdir("final_output", 0777)
	cod, err := os.Create("./final_output/client_owner_domains.txt")
	check(err)
	defer cod.Close()
	sd, err := os.Create("./final_output/suspicious_domains.txt")
	check(err)
	defer sd.Close()

	// sort domains
	var match bool
	var score string
	var w1 *bufio.Writer
	var w2 *bufio.Writer
	var comparedFilePath string
	for _, file := range files {
		if !file.IsDir(){
			comparedFilePath = "./output/" + file.Name()
			match, score = ssdeepTest(clientRefDomainFile, comparedFilePath)
			if match {
				// if match puts "domain matches with domain (score)" in client_owner_domains.txt"
				cod.Sync()
				w1 = bufio.NewWriter(cod)
				_, err = w1.WriteString(cleanString(comparedFilePath) + " matches with " + cleanString(clientRefDomainFile) + " (" + score + ")")
				check(err)
				w1.Flush()
			} else {
				// if doesn't match puts "domain doesn't match with domain (score) - Require more investigation" in suspicious_domains.txt
				sd.Sync()
                                w2 = bufio.NewWriter(sd)
                                _, err = w2.WriteString(cleanString(comparedFilePath) + " doesn't match with " + cleanString(clientRefDomainFile) + " (" + score + ") - Require more investigation")
                                check(err)
                                w2.Flush()
			}
		}
	}

	// Close opened files before exit
	cod.Close()
	sd.Close()
}
