package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Usage = usage

	flag.Parse()

	filters := flag.Args()
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		line = strings.TrimSuffix(line, "\n")


		entry := ParseLTSV(line)
		if len(filters) > 0 {
			entry = entry.Filter(filters)
		}

		fmt.Print(entry, "\n")
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: lq [filter keys...]\n")
	flag.PrintDefaults()
}
