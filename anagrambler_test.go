package anagrambler_test

import (
	"testing"

	"io/ioutil"
	"strings"

	"github.com/RyanEdwardHall/anagrambler"
)

type dataItem struct {
	dict     string
	input    string
	filter   string
	anagrams int
}

// Test data and test fixtures
var (
	testData = []dataItem{
		{"go-dict.txt", "honorificabilitudinitatibus", "", 7214},
		{"go-dict.txt", "honorificabilitudinitatibus", "bus", 29},
		{"go-dict.txt", "pneumonoultramicroscopicsilicovolcanoconiosis", "", 22090},
		{"go-dict.txt", "pneumonoultramicroscopicsilicovolcanoconiosis", "ultra", 24},
		{"go-dict.txt", "Lopadotemachoselachogaleokranioleipsanodrimhypotrimmatosilphioparaomelitokatakechymenokichlepikossyphophattoperisteralektryonoptekephalliokigklopeleiolagoiosiraiobaphetraganopterygon", "", 98616},
		{"go-dict.txt", "Lopadotemachoselachogaleokranioleipsanodrimhypotrimmatosilphioparaomelitokatakechymenokichlepikossyphophattoperisteralektryonoptekephalliokigklopeleiolagoiosiraiobaphetraganopterygon", "pet", 318},
	}
	testTrie *anagrambler.Trie
)

func init() {
	var err error
	testTrie, err = anagrambler.Open(testData[0].dict)

	if err != nil {
		panic(err)
	}
}

func testAnagramCount(t *testing.T, d dataItem) {
	results := testTrie.Search(d.input, d.filter)

	if len(results) == d.anagrams {
		t.Logf("Success: found all %d expected anagrams for '%s' with filter '%s'\n", d.anagrams, d.input, d.filter)
	} else {
		t.Error("Expected", d.anagrams, "words, got ", len(results))
	}
}

func benchmarkOpen(b *testing.B, dictPath string) {
	for counter := 0; counter < b.N; counter++ {
		_, err := anagrambler.Open(dictPath)
		if err != nil {
			panic(err)
		}
	}
}

func benchmarkFillTrie(b *testing.B, dictPath string) {
	data, err := ioutil.ReadFile(dictPath)

	if err != nil {
		b.Error("Could not load dictionary", dictPath, err)
	}

	words := strings.Split(string(data), "\n")
	words = words[:len(words)-1]

	for counter := 0; counter < b.N; counter++ {
		trie := anagrambler.NewTrie()

		for _, word := range words {
			trie.Add(word)
		}
	}
}

func benchmarkSearch(b *testing.B, d dataItem) {
	for counter := 0; counter < b.N; counter++ {
		testTrie.Search(d.input, d.filter)
	}
}

// Tests for verifying correct output (i.e. correct number of anagrams for given input).
func TestAnagramCountShort(t *testing.T)          { testAnagramCount(t, testData[0]) }
func TestAnagramCountShortFiltered(t *testing.T)  { testAnagramCount(t, testData[1]) }
func TestAnagramCountMedium(t *testing.T)         { testAnagramCount(t, testData[2]) }
func TestAnagramCountMediumFiltered(t *testing.T) { testAnagramCount(t, testData[3]) }
func TestAngaramCountLong(t *testing.T)           { testAnagramCount(t, testData[4]) }
func TestAngaramCountLongFilter(t *testing.T)     { testAnagramCount(t, testData[5]) }

// Benchmarks for loading a dictionary file into a trie.
func BenchmarkFillTrie(b *testing.B) { benchmarkFillTrie(b, testData[0].dict) }

func BenchmarkOpen(b *testing.B) { benchmarkOpen(b, testData[0].dict) }

// Benchmarks for searching for anagrams.
func BenchmarkSearchShort(b *testing.B)          { benchmarkSearch(b, testData[0]) }
func BenchmarkSearchShortFiltered(b *testing.B)  { benchmarkSearch(b, testData[1]) }
func BenchmarkSearchMedium(b *testing.B)         { benchmarkSearch(b, testData[2]) }
func BenchmarkSearchMediumFiltered(b *testing.B) { benchmarkSearch(b, testData[3]) }
func BenchmarkSearchLong(b *testing.B)           { benchmarkSearch(b, testData[4]) }
func BenchmarkSearchLongFiltered(b *testing.B)   { benchmarkSearch(b, testData[5]) }
