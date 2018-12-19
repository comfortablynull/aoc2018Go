package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

const R = 32
const A = 'A'

type SkipBuffer struct {
	*bufio.Reader
	skip rune
}

func NewSkipBuffer(r io.Reader, s rune) *SkipBuffer {
	return &SkipBuffer{Reader: bufio.NewReader(r), skip: s}
}

func (b *SkipBuffer) ReadRune() (rune, int, error) {
	for {
		r, i, err := b.Reader.ReadRune()
		if err != nil {
			return r, i, err
		}
		if r == b.skip || r+R == b.skip {
			continue
		}
		return r, i, err
	}
}

func react(buff io.RuneReader) []rune {
	var runes []rune
	for {
		r, _, err := buff.ReadRune()
		if err != nil {
			break
		}
		i := len(runes) - 1
		if len(runes) == 0 || math.Abs(float64(r-runes[i])) != R {
			runes = append(runes, r)
		} else {
			runes = runes[:i]
		}
	}
	return runes
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	runes := react(bufio.NewReader(f))
	str := string(runes)
	fmt.Println("Reacted:", len(runes))
	seen := make(map[rune]int)
	lowest := -1
	for _, v := range runes {
		m := v
		if v < A {
			m += R
		}
		if _, ok := seen[m]; ok {
			continue
		}
		seen[m] = len(react(NewSkipBuffer(strings.NewReader(str), m)))
		if seen[m] < lowest || lowest == -1 {
			lowest = seen[m]
		}
	}
	fmt.Println("Lowest", lowest)
}
