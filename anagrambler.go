package anagrambler

import (
	"io/ioutil"
	"sort"
	"strings"
)

func sortedLower(w string) string {
	w = strings.ToLower(w)
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

type Node struct {
	Words    []string
	Children map[rune]*Node
}

func NewNode() *Node {
	return &Node{
		Words:    make([]string, 0, 1),
		Children: make(map[rune]*Node),
	}
}

func LoadDict(root *Node, filepath string) {
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	words := strings.Split(string(data), "\n")
	words = words[:len(words)-1]

	for _, word := range words {
		AddWord(root, word)
	}
}

func AddWord(root *Node, word string) {
	path := root

	for _, letter := range sortedLower(word) {
		if path.Children[letter] == nil {
			path.Children[letter] = NewNode()
		}
		path = path.Children[letter]
	}
	path.Words = append(path.Words, word)
}

func Search(root *Node, word string) map[*Node]bool {
	results := make(map[*Node]bool)

	search(sortedLower(word), root, results)

	return results
}

func search(postfix string, path *Node, results map[*Node]bool) {
	if len(path.Words) > 0 {
		if !results[path] {
			results[path] = true
		} else {
			return
		}
	}

	searched_runes := make(map[rune]bool)

	for i, letter := range postfix {
		_, nodeExists := path.Children[letter]
		if nodeExists && !searched_runes[letter] {
			search(postfix[i+1:], path.Children[letter], results)

			searched_runes[letter] = true
		}
	}
}
