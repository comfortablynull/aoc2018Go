package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var Offsets = [5]int{-2, -1, 0, 1, 2}

const (
	InitialStatePrefix = "initial state:"
	RuleSep            = " => "
	Full               = '#'
	Empty              = '.'
	SimulateGeneration = 50000000000
	MaxGeneration      = 1000
)

type State map[int]struct{}

func (s State) Sum() int {
	sum := 0
	for k := range s {
		sum += k
	}
	return sum
}
func (s State) String() string {
	start, end := s.Range()
	var str []rune
	for i := start; i < end; i++ {
		if _, ok := s[i]; ok {
			str = append(str, Full)
		} else {
			str = append(str, Empty)
		}
	}
	return string(str)
}

func (s State) Range() (int, int) {
	min, max := math.MaxInt64, math.MinInt64
	for k := range s {
		if k > max {
			max = k
		}
		if k < min {
			min = k
		}
	}
	// adds space so we can evaluate possible matches to the left of the first full pot and the right of the last post
	return min - 3, max + 4
}

func read(in string) (State, map[string]rune) {
	f, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	str := scanner.Text()
	str = strings.TrimSpace(strings.TrimLeft(str, InitialStatePrefix))
	state := make(map[int]struct{})
	for k, v := range str {
		if v == '#' {
			state[k] = struct{}{}
		}
	}
	scanner.Scan()
	rules := map[string]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		ra := strings.Split(line, RuleSep)
		rules[ra[0]] = rune(ra[1][0])
	}
	return state, rules
}

func main() {
	state, rules := read("input.txt")
	seen := map[string]struct{}{}
	for generation := 1; generation < MaxGeneration; generation++ {
		start, end := state.Range()
		newstate := make(map[int]struct{})
		for i := start; i < end; i++ {
			pattern := []rune(strings.Repeat(".", len(Offsets)))
			for k, v := range Offsets {
				if _, ok := state[v+i]; ok {
					pattern[k] = Full
				}
			}
			if v, ok := rules[string(pattern)]; ok && v == Full {
				newstate[i] = struct{}{}
			}
		}
		state = newstate
		if generation == 20 {
			fmt.Println("20:", state.Sum())
		}
		pattern := strings.Trim(state.String(), ".")
		// check to see if the pattern has repeated. If it has assumed that it has stabilized
		if _, ok := seen[pattern]; ok {
			// since the pattern is stable it is going to move to the right
			// in order to calculate the shifted sum we take the difference between the current generation
			// and the desired generation add that to the position of the pot in order to get it's location.
			offset := SimulateGeneration - generation
			sum := 0
			for k := range state {
				sum += k + offset
			}
			fmt.Println(SimulateGeneration, ":", sum)
			return
		}
		seen[pattern] = struct{}{}
	}
	fmt.Println("Could simulate")
}
