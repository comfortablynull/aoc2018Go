package day6

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X  int
	Y  int
	A  int
	UC bool
}

type Scanner struct {
	*bufio.Scanner
}

func NewScanner(r io.Reader) *Scanner {
	buff := bufio.NewScanner(r)
	return &Scanner{buff}
}

func (s *Scanner) Point() (*Point, error) {
	str := s.Text()
	if str == "" {
		return nil, errors.New("invalid")
	}
	xy := strings.Split(str, ", ")
	x, err := strconv.ParseInt(xy[0], 10, 64)
	if err != nil {
		return nil, err
	}
	y, err := strconv.ParseInt(xy[1], 10, 64)
	if err != nil {
		return nil, err
	}
	return &Point{X: int(x), Y: int(y)}, nil
}

type Points []*Point

func (p Points) Lowest(x, y, w, h int) (*Point, int) {
	var point *Point
	td := 0
	lowest := -1
	for _, v := range p {
		d := int(math.Abs(float64(x-v.X)) + math.Abs(float64(y-v.Y)))
		if lowest == -1 || d < lowest {
			lowest = d
			point = v
		} else if lowest == d {
			point = nil
		}
		td += d
	}
	if point != nil {
		if x == 0 || y == 0 {
			point.UC = true
		}
		if x == w || y == h {
			point.UC = true
		}
		point.A++
	}
	return point, td
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var points Points
	w, h := 0, 0
	scanner := NewScanner(f)
	for scanner.Scan() {
		p, err := scanner.Point()
		if err != nil {
			break
		}
		points = append(points, p)
		if p.X > w {
			w = p.X
		}
		if p.Y > h {
			h = p.Y
		}
	}
	area := points[0]
	total := 0
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			lowest, td := points.Lowest(x, y, w, h)
			if lowest != nil && !lowest.UC && lowest.A > area.A {
				area = lowest
			}
			if td < 10000{
				total += 1
			}
		}
	}
	fmt.Println("largest", area.A)
	fmt.Println("total", total)
}
