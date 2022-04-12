package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a booelan flag -l to count lines instead of words
	words := flag.Bool("w", false, "Count words")
	bytes := flag.Bool("b", false, "Count bytes")

	// Parsing the flag provided by the user
	flag.Parse()

	// Calling the count function to count the number of words (or lines)
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *words, *bytes))

}

func count(r io.Reader, countWords bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	switch {
	case countWords:
		scanner.Split(bufio.ScanWords)
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
