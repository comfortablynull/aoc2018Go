package main

import (
	"fmt"
)

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func (n *Node) Insert(nn *Node) {
	nn.Next = n.Next
	n.Next.Previous = nn
	n.Next = nn
	nn.Previous = n
}

func (n *Node) Pop(b int) *Node {
	if b > 0 {
		b--
		return n.Previous.Pop(b)
	}
	n.Previous.Next = n.Next
	n.Next.Previous = n.Previous
	return n
}

func main() {
	mx := 71052 * 100
	max := (mx / 23) * 23
	node := &Node{}
	node.Next = node
	node.Previous = node
	numPlayers := 419
	players := make(map[int]int)
	highscore := 0
	for i := 1; i <= max; i++ {
		n := &Node{Value: i}
		if i%23 == 0 {
			node = node.Pop(7)
			k := (i - 1) % numPlayers
			players[k] += i
			players[k] += node.Value
			if players[k] > highscore {
				highscore = players[k]
			}
			node = node.Next
			continue
		}
		node.Next.Insert(n)
		node = n
	}
	fmt.Println(highscore)
}