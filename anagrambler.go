package anagrambler

import (
	"bytes"
	"io/ioutil"
	"sort"
	"unicode/utf8"
)

// Types

type Trie struct {
	root *node
}

type node struct {
	Words    [][]byte
	Children map[rune]*node
}

// sortBytes is necessary because sort.Sort requires a named type
type sortBytes []byte

// Exported functions

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func Open(filepath string) (*Trie, error) {
	t := NewTrie()

	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	words := bytes.Split(data, []byte("\n"))
	words = words[:len(words)-1]

	for _, word := range words {
		sortedWord := sortWord(bytes.ToLower(word))
		t.add(word, sortedWord)
	}

	return t, nil
}

// Exported Methods

func (t *Trie) Add(word string) {

	sortedWord := sortWord(bytes.ToLower([]byte(word)))

	t.add([]byte(word), sortedWord)
}


func (t *Trie) Search(text string, filter string) []string {
	results := make(map[*node]bool)

	sortedText, sortedFilter := sortWord(bytes.ToLower([]byte(text))), sortWord(bytes.ToLower([]byte(filter)))

	search(t.root, sortedText, sortedFilter, results)

	filteredResults := make([]string, 0)

	for node := range results {
		for _, word := range node.Words {
			if bytes.Contains(word, []byte(filter)) {
				filteredResults = append(filteredResults, string(word))
			}
		}
	}

	return filteredResults
}

func (t *Trie) add(word, sortedWord []byte) {
	path := t.root

	for i, w := 0, 0; i < len(sortedWord); i += w {
		r, width := utf8.DecodeRune(sortedWord[i:])

		if path.Children[r] == nil {
			path.Children[r] = newNode()
		}
		path = path.Children[r]

		w = width
	}

	path.Words = append(path.Words, word)
}

// Unexported Functions

func newNode() *node {
	return &node{
		Words:    make([][]byte, 0, 1),
		Children: make(map[rune]*node),
	}
}

func (s sortBytes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
    return len(s)
}

func sortWord(w []byte) []byte {
	sort.Sort(sortBytes(w))
	return w
}

func search(n *node, text []byte, filter []byte, results map[*node]bool) {
	// Record any words stored at this node
	// Only record acronyms after the filter has been satisfied
	if len(filter) == 0 && len(n.Words) != 0 {
		if !results[n] {
			// Add this node's acronyms to the results
			results[n] = true
		} else {
			// We've already traversed this node, so stop searching it
			return
		}
	}

	// Keep track of which runes we've searched
	searched_runes := make(map[rune]bool)

	for i, w := 0, 0; i < len(text); i += w {
		r, width := utf8.DecodeRune(text[i:])
		w = width

		// Skip any runes that we don't have nodes for
		// or that we've already searched for (i.e. duplicate runes)
		if n.Children[r] == nil || searched_runes[r] == true {
			continue
		}

		if len(filter) != 0 {
			fRune, fWidth := utf8.DecodeRune(filter)

			if r == fRune {
				// This letter satisfies the next rune in the filter, so we can
				// remove it from the filter
				filter = filter[fWidth:]
			} else if r > fRune {
				// The remaining letters in the text are all greater than the next
				// required filter rune, so none of the remaining substrings will
				// satisfy the filter
				return
			}
		}

		search(n.Children[r], text[i+width:], filter, results)

		searched_runes[r] = true
	}
}
