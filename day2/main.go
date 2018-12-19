package main

import (
	"aoc2018Go/lib"
	"log"
	"os"
	"sort"
)

func checksum(str string) int {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	l, s := len(r), len(r)-1
	sum := 0
	c := 0
	for i := 0; i < l; i++ {
		if i != s && r[i] == r[i+1] {
			c++
			continue
		}
		if c == 1 && sum != 1 {
			sum += 1
		} else if c == 2 && sum != 2 {
			sum += 2
		}
		if sum == 3 {
			break
		}
		c = 0
	}
	return sum
}

func sum(scanner *lib.AdventScanner) (int, error) {
	if err := scanner.Rewind(); err != nil {
		return 0, err
	}
	a, b := 0, 0
	for scanner.Scan() {
		s := checksum(scanner.Text())
		a += s % 2
		b += s / 2
	}
	return a * b, nil
}

func diff(scanner *lib.AdventScanner) (string, error) {
	if err := scanner.Rewind(); err != nil {
		return "", err
	}
	var s [][]byte
	for scanner.Scan() {
		s = append(s, scanner.Bytes())
	}
	for k, v := range s {
		for _, vv := range s[k:] {
			d := 0
			i := 0
			for kk, c := range vv {
				if c != v[kk] {
					i = kk
					d++
				}
				if d > 1 {
					break
				}
			}
			if d == 1 {
				return string(append(v[:i], v[i+1:]...)), nil
			}
		}
	}
	return "", nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := lib.NewAdventScanner(f)
	s, err := sum(scanner)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("sum:", s)
	c, err := diff(scanner)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("closest:", c)
}
