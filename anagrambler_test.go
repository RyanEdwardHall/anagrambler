package anagrambler_test

import (
	"testing"

	"github.com/RyanEdwardHall/anagrambler"
)

func TestKnownOutput(t *testing.T) {
	trie := anagrambler.NewNode()

	anagrambler.LoadDict(trie, "go-dict.txt")

	searchWord := "honorificabilitudinitatibus"

	results := anagrambler.Search(trie, searchWord)

	counter := 0
	for path := range results {
		counter += len(path.Words)
	}
	if counter != 9083 {
		t.Error("Expected 9083 words, got ", counter)
	}
}

func BenchmarkAnagrambler(b *testing.B) {
	trie := anagrambler.NewNode()

	anagrambler.LoadDict(trie, "go-dict.txt")

	searchWord := "Lopadotemachoselachogaleokranioleipsanodrimhypotrimmatosilphioparaomelitokatakechymenokichlepikossyphophattoperisteralektryonoptekephalliokigklopeleiolagoiosiraiobaphetraganopterygon"

	b.ResetTimer()

	for counter := 0; counter < b.N; counter++ {
		anagrambler.Search(trie, searchWord)
	}
}
