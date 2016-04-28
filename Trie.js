'use strict';

var words = require("an-array-of-english-words");

class Trie {
  constructor(words) {
    this.data = {};
    this.insert(words);
  }
  insert(words) {
    words.forEach(word => {
      var englishWord = word;
      word = word.split('').sort().join('');
      var characters = word.split('');
      var path = this.data;
      characters.forEach((character, index) => {
        if (!path.hasOwnProperty(character)) {
          path[character] = {};
        }
        if (index === (word.length - 1)) {
          if (!path[character].hasOwnProperty('_words')) {
            path[character]._words = [];
          }
          path[character]._words.push(englishWord);
        } else {
          path = path[character];
        }
      });
    });
  }
  search(letters) {
    var noDupes = {};
    var self = this;
    letters = letters.split('').sort();
    function searcher(prefix, remainder, path) {
      if (noDupes[prefix] === undefined) {
        noDupes[prefix] = null;
      }
      try {
        path = path[prefix.slice(-1)];
        if (path._words !== undefined) {
          noDupes[prefix] = path._words;
        }
      } catch (e) {
        path = null;
      }

      remainder.forEach((letter, index) => {
        var remainderCopy = remainder.slice();
        var popRemainder = remainderCopy.splice(index, 1);
        var newPrefix = (prefix + popRemainder).split('').sort().join('');
        if (noDupes[newPrefix] === undefined) {
          return searcher(newPrefix, remainderCopy, path);
        }
      });
    }

    while (letters.length > 1) {
      searcher(letters.shift(), letters, self.data);
    }

    Object.keys(noDupes).forEach(key => {
      if (noDupes[key] !== null) {
        console.log(noDupes[key]);
      }
    });
  }
}

var TrieEnglishDic = new Trie(words);
TrieEnglishDic.search('test');
