package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {

	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true

}

// Search will take in a word and returns true, is that word included in the trie
func (t *Trie) Search(w string) bool {

	currentNode := t.root

	for _, c := range w {
		charIndex := c - 'a'
		if currentNode.children[charIndex] == nil {
			break
		}
		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd

}

func main() {

	myTrie := InitTrie()
	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	myTrie.Insert("aragrn")
	fmt.Println(myTrie.Search("orego"))

}
