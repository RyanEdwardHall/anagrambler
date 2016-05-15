package main

import "testing"

func BenchmarkAnagrambler(b *testing.B) {
	root := &node{
		words:    make([]string, 0, 1),
		children: make(map[string]*node),
	}
	path := root
	LoadTrie(path, root)
	searchWord := SortString("honorificabilitudinitatibus")
	results := make(map[*node]bool)

	b.ResetTimer()
	for counter := 0; counter < b.N; counter++ {
		for i := range searchWord {
			_, nodeExists := path.children[searchWord[i:i+1]]
			if nodeExists {
				search(searchWord[i:i+1], searchWord[i+1:], path.children[searchWord[i:i+1]], results)
			}
		}
	}
}
