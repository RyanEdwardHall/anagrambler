package main

import (
    "io/ioutil"
    "fmt"
    "strings"
    "sort"
    "os"
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
    words []string
    children map[string]*node
}

func search (prefix string, postfix string, path *node, results map[string]string) {
  if len(path.words) > 0 {
    fullPath := SortString(path.words[0])
    _, exists := results[fullPath]
    if !exists {
      results[fullPath] = strings.Join(path.words, ",")
    }
  }
  for i, letter := range postfix {
    _, nodeExists := path.children[string(letter)]
    if nodeExists {
      search(string(letter), postfix[i+1:], path.children[string(letter)], results)
    }
  }
}

func main() {
  // load dictionary
  data, err := ioutil.ReadFile("go-dict.txt")
  // data, err := ioutil.ReadFile("testwords.txt")
  check(err)

  // initialize the data structure
  root := &node{
    words: make([]string, 0, 50),
    children: make(map[string]*node),
  }

  path := root
  words := strings.Split(string(data), "\n")
  words = words[:len(words) - 1]
  for i := 0; i < len(words); i++ {
    englishWord := words[i]
    words[i] = strings.ToLower(words[i])
    words[i] = SortString(words[i])
    path = root
    for x := 0; x < len(words[i]); x++ {
      letter := string(words[i][x])
      _, exists := path.children[letter]; 
      if !exists {
        path.children[letter] = &node{
          words: make([]string, 0, 50),
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
  
  //search trie (adnor has lots of anagrams)
  path = root
  searchWord := SortString(os.Args[1])
  results := make(map[string]string, 100)
  for i := range searchWord {
    _, nodeExists := path.children[searchWord[i:i+1]]
    if nodeExists {
      search(searchWord[i:i+1], searchWord[i+1:], path.children[searchWord[i:i+1]], results)
    }
  }

  for i:= range results {
    fmt.Println(results[i])    
  }
}
