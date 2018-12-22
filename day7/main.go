package day7

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

const None = rune(91)

type Step struct {
	C     rune
	B     []*Step //before
	ready bool
	Ran   bool
	Done  bool
	T     int
}

func NewStep(c rune) *Step {
	return &Step{C: c, T: int(c - 4)}
}

func (s *Step) Ready() bool {
	for _, v := range s.B {
		if !v.Ran {
			return false
		}
	}
	return true
}

func MakeSteps(r io.Reader) []*Step {
	stepMap := make(map[rune]*Step)
	var steps []*Step
	buff := bufio.NewScanner(r)
	for buff.Scan() {
		text := buff.Text()
		b, a := rune(text[5]), rune(text[36])
		as, ok := stepMap[a]
		if !ok {
			as = NewStep(a)
			stepMap[a] = as
			steps = append(steps, as)
		}
		bs, ok := stepMap[b]
		if !ok {
			bs = NewStep(b)
			stepMap[b] = bs
			steps = append(steps, bs)
		}
		as.B = append(as.B, bs)
	}
	sort.Slice(steps, func(i, j int) bool {
		return steps[i].C < steps[j].C
	})
	return steps
}

type Nexter struct {
	steps []*Step
}

func (n *Nexter) Done() bool {
	return len(n.steps) == 0
}
func NewNexter(s []*Step) *Nexter {
	n := &Nexter{steps: make([]*Step, len(s))}
	copy(n.steps, s)
	for _, v := range n.steps {
		v.Ran = false
		v.Done = false
		v.ready = false
	}
	return n

}
func (n *Nexter) Next() *Step {
	var step *Step
	for k, v := range n.steps {
		if !v.Ready() {
			continue
		}
		step = v
		n.steps = append(n.steps[:k], n.steps[k+1:]...)
		break
	}
	return step
}

type Worker struct {
	s *Step
}

func (w *Worker) Set(s *Step) {
	w.s = s
}

func (w *Worker) Step() bool {
	if w.s == nil {
		return true
	}
	if w.s.T == 0 {
		w.s.Ran = true
		return true
	}
	w.s.T--
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	steps := MakeSteps(f)
	var r []rune
	nexterA := NewNexter(steps)
	for !nexterA.Done() {
		p := nexterA.Next()
		p.Ran = true
		r = append(r, p.C)
	}
	fmt.Println(string(r))
	s := 0
	var workers []*Worker
	numWorkers := 5
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, &Worker{})
	}
	nexter := NewNexter(steps)
	drained := false
	for !nexter.Done() || !drained {
		drained = true
		for _, v := range workers {
			if !v.Step() {
				drained = false
			} else if j := nexter.Next(); j != nil {
				v.Set(j)
				v.Step()
				drained = false
			}
		}
		if !drained {
			s++
		}
	}
	fmt.Println(s)
	// NOT ADLXVJEFKBWCQUNGORTMYSIHPZ
	//	fmt.Println(string(start.Run(nil)))
}
