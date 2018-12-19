package main

import (
	"aoc2018Go/lib"
	"log"
	"os"
)

func total(scanner *lib.AdventScanner) (int64, error) {
	if err := scanner.Rewind(); err != nil {
		return 0, err
	}
	t := int64(0)
	for scanner.Scan() {
		i, err := scanner.Int()
		if err != nil {
			return t, err
		}
		t += i
	}
	return t, nil
}

func twice(scanner *lib.AdventScanner) (int64, error) {
	state := int64(0)
	seen := make(map[int64]bool)
	for {
		if err := scanner.Rewind(); err != nil {
			return 0, err
		}
		for scanner.Scan() {
			seen[state] = true
			i, err := scanner.Int()
			if err != nil {
				return 0, err
			}
			state += i
			if _, ok := seen[state]; ok {
				return state, nil
			}
		}
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := lib.NewAdventScanner(f)
	t, err := total(scanner)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total:", t)
	tt, err := twice(scanner)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("twice:", tt)
}
