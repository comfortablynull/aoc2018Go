package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const position, velocity = "position=", "velocity="
const start, sep, end = '<', ',', '>'

type Pos struct {
	X, Y int
}
type Point struct {
	P *Pos
	V Pos
}

func parseInt(str string) int {
	t, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Panic(err)
	}
	return int(t)
}

func read(str, prefix string) Pos {
	l := len(prefix)
	for i := 0; i < len(str); i += 1 {
		if str[i:i+l] == prefix {
			str = str[i+l:]
			break
		}
	}
	var x, y int
	ref := 0
	for i := 0; i < len(str); i++ {
		if str[i] == start {
			ref = i
		} else if str[i] == sep {
			x = parseInt(str[ref+1 : i])
			ref = i
		} else if str[i] == end {
			y = parseInt(str[ref+1 : i])
			return Pos{
				X: x,
				Y: y,
			}
		} else if str[i] == ' ' {
			ref = i
		}
	}
	panic("couldn't parse")
}

func parse(str string) Point {
	p := read(str, position)
	v := read(str, velocity)
	return Point{
		P: &p,
		V: v,
	}
}

func load(path string) []Point {
	var p []Point
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		p = append(p, parse(line))
	}
	return p
}

type Graph struct {
	graph                  map[int]map[int]Pos
	maxX, minX, maxY, minY int
}

func NewGraph() *Graph {
	return &Graph{graph: map[int]map[int]Pos{}}
}

func (g *Graph) Add(p Pos) {
	if _, ok := g.graph[p.Y]; !ok {
		g.graph[p.Y] = map[int]Pos{}
	}
	if g.maxX < p.X {
		g.maxX = p.X
	}
	if g.minX > p.X {
		g.minX = p.X
	}
	if g.minY > p.Y {
		g.minY = p.Y
	}
	if g.maxY < p.Y {
		g.maxY = p.Y
	}
	g.graph[p.Y][p.X] = p
}

func (g *Graph) Area() int {
	return (g.maxX - g.minX + 1) + (g.maxY + g.minY + 1)
}

func (g *Graph) Print() {
	var order []int
	for k := range g.graph {
		order = append(order, k)
	}
	sort.Ints(order)
	max := g.maxX - g.minX + 1
	for _, key := range order {
		row := g.graph[key]
		str := make([]rune, max)
		for k := range str {
			if _, ok := row[k-g.minX]; ok {
				str[k] = '#'
			} else {
				str[k] = '.'
			}
		}
		fmt.Println(string(str))
	}
}

func main() {
	points := load("input.txt")
	var previous *Graph
	for i := 0; ; i++ {
		g := NewGraph()
		for _, v := range points {
			p := v.P
			vv := v.V
			p.X += vv.X
			p.Y += vv.Y
			g.Add(*p)
		}
		if previous != nil && previous.Area() < g.Area() {
			fmt.Println("Seconds", i)
			previous.Print()
			return
		}
		previous = g
	}
}
