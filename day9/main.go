package day9

import (
	"fmt"
	"math"
)

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func main() {
	mx := 71052*100
	max := (mx / 23) * 23
	node := &Node{}
	node.Next = node
	node.Previous = node
	numPlayers := 419
	players := make([]int, numPlayers)
	highscore := 0
	for i := 1; i <= max; i++ {
		n := &Node{Value: i}
		if i%23 == 0 {
			for i := 0; i < 7; i++ {
				node = node.Previous
			}
			k := (i - 1) % numPlayers
			players[k] += i
			players[k] += node.Value
			if players[k] > highscore {
				highscore = players[k]
			}
			node.Previous.Next = node.Next
			node.Next.Previous = node.Previous
			node = node.Next
			continue
		}
		in := node.Next
		n.Next = in.Next
		in.Next.Previous = n
		in.Next = n
		n.Previous = in
		node = n
	}
	fmt.Println(highscore)
}
