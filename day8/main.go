package day8

import (
	"aoc2018Go/lib"
	"bufio"
	"fmt"
	"log"
	"os"
)

type Node struct {
	children []*Node
	metadata []int
}

func NewNode(c, m int) *Node {
	return &Node{children: make([]*Node, c, c), metadata: make([]int, m, m)}
}
func MakeNode(scanner *lib.AdventScanner) *Node {
	scanner.Scan()
	cc, _ := scanner.Int()
	scanner.Scan()
	mm, _ := scanner.Int()
	return NewNode(int(cc), int(mm)).ReadChildren(scanner).ReadMeta(scanner)
}

func (n *Node) ReadChildren(scanner *lib.AdventScanner) *Node {
	for i := 0; i < len(n.children); i ++ {
		n.children[i] = MakeNode(scanner)
	}
	return n
}

func (n *Node) ReadMeta(scanner *lib.AdventScanner) *Node {
	for i := 0; i < len(n.metadata) && scanner.Scan(); i++ {
		v, err := scanner.Int()
		if err != nil {
			continue
		}
		n.metadata[i] = int(v)
	}
	return n
}

func (n *Node) Sum() int {
	t := 0
	for _, v := range n.metadata {
		t += v
	}
	for _, v := range n.children {
		t += v.Sum()
	}
	return t
}

func (n *Node) WeirdSum() int {
	l := len(n.children)
	if l == 0 {
		return n.Sum()
	}
	t := 0
	for _, v := range n.metadata {
		k := v - 1
		if k >= l {
			continue
		}
		t += n.children[k].WeirdSum()
	}
	return t
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := lib.NewAdventScanner(f)
	scanner.Split(bufio.ScanWords)
	n := MakeNode(scanner)
	fmt.Println(n.Sum())
	fmt.Println(n.WeirdSum())
}
