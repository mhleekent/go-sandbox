package main

import (
	// "bytes"
    "os"
    "bufio"
	"io"
    "fmt"
    "strings"
    // "log"
)

const (
    LONGEST_MATCH = iota
    SHORTEST_MATCH
    MATCH_ALL
)

func populateTrie(trie *Trie, file string) {
    f, err := os.Open(file)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    rd := bufio.NewReader(f)
    
    for {
        line, err := rd.ReadString('\n')
        if err == io.EOF {
            break
        } else {
            trie.AddWord(strings.TrimRight(line, "\r\n"))
        }
    }
}

func traverse(node *TrieNode, level int) {
    // log.Println(level, "> Letter:", node.Letter)
    for _, node := range(node.children) {
        traverse(node, level+1)
    }
}

func foundWord(words []string) {
    // var buffer bytes.Buffer
    fmt.Printf("Found word: ")
    for _, word := range(words) {
        fmt.Printf("%s ", word)
    }
    fmt.Printf("\n")
}

func main(){
    if len(os.Args) < 2 || len(os.Args) > 3 {
        fmt.Printf("Usage: %s <dictionary file> [input text]\n", os.Args[0])
        os.Exit(0)
    }
    
    dicFilePath := os.Args[1]
    var reader *bufio.Reader 
    if len(os.Args) == 3 {
        f, err := os.Open(os.Args[2])
        if err != nil {
            panic(err)
        }
        defer f.Close()
        reader = bufio.NewReader(f)
    } else {
        reader = bufio.NewReader(os.Stdin)
    }
    
    trie := NewTrie()
    populateTrie(trie, dicFilePath)
    traverse(trie.Root, 0)
    
    for {
        
        if len(os.Args) == 2 {
            fmt.Print("Enter words: ")
        }
        text, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }
        
        words := strings.Split(strings.TrimRight(text, "\r\n"), " ")
        
        start := 0
        var node = trie.Root
        var ok = false
        
        for i, word := range(words) {
            // log.Println("Word: ", word)
            node, ok = trie.FindWord(word, node)
            if ok == false {
                start = i + 1
                node = trie.Root
            } else if node.IsLeaf() == true { 
                foundWord(words[start:i+1])
                start = i+1
                node = trie.Root
            }
        } 
    }
}