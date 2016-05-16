package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/RyanEdwardHall/anagrambler"
)

func main() {
	trie := anagrambler.NewNode()

	anagrambler.LoadDict(trie, "go-dict.txt")

	searchWord := os.Args[1]

	results := anagrambler.Search(trie, searchWord)

	count := 0

	for path := range results {
		fmt.Println(path.Words)
		count += len(path.Words)
	}

	fmt.Println("Number of anagrams:", count)
}
