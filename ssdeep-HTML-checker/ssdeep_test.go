package main

import (
	"github.com/glaslos/ssdeep"
	"log"
	"strconv"
)

func ssdeepTest(fnameRef string, fnameComp string) (bool,string) {
	// false = don't try to get a ssdeep hash if filesize is lesser than 4096 bytes
	// true = bypass the error to have a hash (not recommended)
	ssdeep.Force = false
	var result string

	h1, err := ssdeep.FuzzyFilename(fnameRef)
	if err != nil && !ssdeep.Force {
		log.Fatal(err)
	}

	h2, err := ssdeep.FuzzyFilename(fnameComp)
	if err != nil && !ssdeep.Force {
                log.Fatal(err)
        }

	score, err = ssdeep.Distance (h1, h2)
	result, _ = strconv.Itoa(score)
	if score != 0 {
		return true, result
	} else if err != nil {
		log.Fatal(err)
	} else {
		return false, result
	}
}
