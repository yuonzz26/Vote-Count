package main

import (
	"fmt"
	"log"

	"github.com/Vote-Count/votecounts"
)

func main() {
	lines, err := votecounts.GetStrings("votes.txt")

	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
