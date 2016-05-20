package anagrambler_test

import (
	"testing"

	"github.com/RyanEdwardHall/anagrambler"
)

type dataItem struct {
	dict string
	input string
	anagrams int
}

var testData = []dataItem {
	{"go-dict.txt", "honorificabilitudinitatibus", 9083},
	{"go-dict.txt", "Lopadotemachoselachogaleokranioleipsanodrimhypotrimmatosilphioparaomelitokatakechymenokichlepikossyphophattoperisteralektryonoptekephalliokigklopeleiolagoiosiraiobaphetraganopterygon", 112436},
}

func testAnagramCount(t *testing.T, d dataItem) {
	trie := anagrambler.NewNode()

	anagrambler.LoadDict(trie, d.dict)

	searchWord := d.input

	results := anagrambler.Search(trie, searchWord, "")

	if len(results) == d.anagrams {
		t.Logf("Success: found all %d expected anagrams for '%s'\n", d.anagrams, d.input)
	} else {
		t.Error("Expected", d.anagrams, "words, got ", len(results))
	}
}

func benchmarkSearch(b *testing.B, d dataItem) {
	trie := anagrambler.NewNode()

	anagrambler.LoadDict(trie, d.dict)

	searchWord := d.input

	b.ResetTimer()

	for counter := 0; counter < b.N; counter++ {
		anagrambler.Search(trie, searchWord, "")
	}
}

func TestAnagramCountShort(t *testing.T) { testAnagramCount(t, testData[0]) }
func TestAngaramCountLong(t *testing.T) { testAnagramCount(t, testData[1]) }

func BenchmarkSearchShort(b *testing.B) { benchmarkSearch(b, testData[0]) }
func BenchmarkSearchLong(b *testing.B) { benchmarkSearch(b, testData[1]) }
