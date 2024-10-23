package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines, bytes bool) int {
	scanner := bufio.NewScanner(r)
	switch {
	case bytes:
		scanner.Split(bufio.ScanBytes)
	case !countLines:
		scanner.Split(bufio.ScanWords)
	default:
		scanner.Split(bufio.ScanLines)
	}
	counter := 0
	for scanner.Scan() {
		counter++
	}
	return counter
}
