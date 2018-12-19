package main

import (
	"aoc2018Go/lib"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	point map[string][]int
	multi map[string]bool
	runs  map[int]bool
}

func NewGraph() *Graph {
	return &Graph{point: make(map[string][]int), multi: make(map[string]bool), runs: make(map[int]bool)}
}

func (g *Graph) Plot(r, x, yS, yE int) {
	g.runs[r] = true
	for y := yS; y < yE; y++ {
		s := fmt.Sprintf("%v,%v", x, y)
		g.point[s] = append(g.point[s], r)
		if len(g.point[s]) > 1 {
			g.multi[s] = true
		}
	}
}

func (g *Graph) Collisions() int {
	return len(g.multi)
}

func (g *Graph) Whole() int {
	for k := range g.multi {
		for _, v := range g.point[k] {
			delete(g.runs, v)
		}
	}
	for k := range g.runs {
		return k
	}
	return -1
}

func parseXY(s string) (int, int, error) {
	point := strings.Split(s, ",")
	x, err := strconv.ParseInt(point[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.ParseInt(point[1][:len(point[1])-1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return int(x), int(y), nil
}

func parseSize(s string) (int, int, error) {
	size := strings.Split(s, "x")
	w, err := strconv.ParseInt(size[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	h, err := strconv.ParseInt(size[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return int(w), int(h), nil

}

func parseRun(s string) (int, error) {
	r, err := strconv.ParseInt(s[1:], 10, 64)
	return int(r), err
}

func plot(cord string, graph *Graph) (int, error) {
	s := strings.Split(cord, " ")
	r, err := parseRun(s[0])
	if err != nil {
		return 0, err
	}
	x, y, err := parseXY(s[2])
	if err != nil {
		return r, err
	}
	w, h, err := parseSize(s[3])
	if err != nil {
		return r, err
	}
	for i := 0; i < w; i++ {
		graph.Plot(r, i+x, y, y+h)
	}
	return r, nil
}

func A(scanner *lib.AdventScanner) (*Graph, error) {
	graph := NewGraph()
	runs := make(map[int]bool)
	for scanner.Scan() {
		r, err := plot(scanner.Text(), graph)
		if err != nil {
			return nil, err
		}
		runs[r] = true
	}
	return graph, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := lib.NewAdventScanner(f)
	graph, err := A(scanner)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Collisions:", graph.Collisions())
	log.Println("Complete:", graph.Whole())
}
