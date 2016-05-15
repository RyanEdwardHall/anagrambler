package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

type node struct {
	words    []string
	children map[string]*node
}

func search(postfix string, path *node, results map[*node]bool) {
	if len(path.words) > 0 {
		if !results[path] {
			results[path] = true
		} else {
			return
		}
	}
	searched_runes := make(map[rune]bool)
	for i, letter := range postfix {
		_, nodeExists := path.children[string(letter)]
		if nodeExists && !searched_runes[letter] {
			search(postfix[i+1:], path.children[string(letter)], results)
		}
		if !searched_runes[letter] {
			searched_runes[letter] = true
		}
	}
}

func LoadTrie(path *node, root *node) {
	data, err := ioutil.ReadFile("go-dict.txt")
	check(err)

	words := strings.Split(string(data), "\n")
	words = words[:len(words)-1]
	for i := 0; i < len(words); i++ {
		englishWord := words[i]
		words[i] = strings.ToLower(words[i])
		words[i] = SortString(words[i])
		path = root
		for x := 0; x < len(words[i]); x++ {
			letter := string(words[i][x])
			_, exists := path.children[letter]
			if !exists {
				path.children[letter] = &node{
					words:    make([]string, 0, 1),
					children: make(map[string]*node),
				}
			}
			if len(words[i]) == (x + 1) {
				path.children[letter].words = append(path.children[letter].words, englishWord)
			} else {
				path = path.children[letter]
			}
		}
	}
}

func main() {
	root := &node{
		words:    make([]string, 0, 1),
		children: make(map[string]*node),
	}
	path := root

	LoadTrie(path, root)

	//search trie (adnor has lots of anagrams)
	path = root
	searchWord := SortString(os.Args[1])
	results := make(map[*node]bool)
	for i := range searchWord {
		_, nodeExists := path.children[searchWord[i:i+1]]
		if nodeExists {
			search(searchWord[i+1:], path.children[searchWord[i:i+1]], results)
		}
	}

	counter := 0
	for path := range results {
		counter += len(path.words)
	}
	fmt.Println(counter)
}
