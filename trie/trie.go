package main

import (
    // "log"
)

type Trie struct {
    Root *TrieNode
} 

func NewTrie() *Trie {
    trie := new(Trie)
    trie.Root = NewTrieNode("")
    return trie
}

func (trie *Trie) AddWord(word string) {
    if len(word) < 1 {
        return 
    }
    node := trie.Root
    for _, c := range word {
        child := node.AddChild(string(c))
        node = child
    }
    node.SetWord(true)
}

func (trie *Trie) FindWord(word string, node *TrieNode) (lastNode *TrieNode, ok bool) {
    for _, c := range word {
        node = node.GetChild(string(c))
        
        if node == nil {
            return new(TrieNode), false
        }
        // log.Println("   this.Letter:", node.Letter, "children:", node.children)
    }
    return node, true
}