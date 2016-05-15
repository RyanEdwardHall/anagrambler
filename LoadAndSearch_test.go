package main

import "testing"

func BenchmarkAnagrambler(b *testing.B) {
	root := &node{
		words:    make([]string, 0, 1),
		children: make(map[rune]*node),
	}
	path := root
	LoadTrie(path, root)
	searchWord := SortString("honorificabilitudinitatibus")
	results := make(map[*node]bool)

	b.ResetTimer()
	for counter := 0; counter < b.N; counter++ {
		for i, letter := range searchWord {
			_, nodeExists := path.children[letter]
			if nodeExists {
				search(searchWord[i+1:], path.children[letter], results)
			}
		}
	}
}
