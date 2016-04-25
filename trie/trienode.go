package main

// import "log"

type TrieNode struct {
    Letter string
    children map[string]*TrieNode
    isWord bool
}

func NewTrieNode(letter string) *TrieNode {
    //node := TrieNode{ letter, nil, false}
    node := new(TrieNode)
    node.Letter = letter
    node.children = make(map[string]*TrieNode)
    node.isWord = false
    return node
}

func (node *TrieNode) AddChild(letter string) *TrieNode {
    child, ok := node.children[letter]
    if ok == true {
        return child
    } else {
        child = NewTrieNode(letter)
        node.children[letter] = child
        return child
    }
}

func (node *TrieNode) GetChild(letter string) *TrieNode {
    child, ok := node.children[letter]
    if ok == true {
        return child
    }
    return nil
}

func (node *TrieNode) SetWord(ok bool) {
    node.isWord = ok
}

func (node *TrieNode) IsWord() (ok bool) {
    return node.isWord
}

func (node *TrieNode) IsLeaf() (ok bool ) {
    return len(node.children) == 0 
}
