package main

import (
	"fmt"
	"os"

	"github.com/RyanEdwardHall/anagrambler"
)

func main() {
	trie := anagrambler.NewTrie()

	anagrambler.LoadDict(trie, "go-dict.txt")

	if len(os.Args) > 1 {
		searchWord := os.Args[1]

		filter := ""

		if len(os.Args) == 3 {
			filter = os.Args[2]
		}

		results := anagrambler.Search(trie, searchWord, filter)

		fmt.Println("Number of anagrams:", len(results))

		for _, anagram := range results {
			fmt.Println(anagram)
		}
	} else {
		fmt.Println("No search string specified")
	}
}
