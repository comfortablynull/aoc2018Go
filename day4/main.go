package main

import (
	"aoc2018Go/lib"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

const TimeLayout = `[2006-01-02 15:04]`

type Guard struct {
	Id      int64
	Sleep   time.Duration
	Freq    map[int]int
	Highest int
}

func NewGuard(id string) *Guard {
	i, _ := strconv.ParseInt(id, 10, 64)
	return &Guard{Id: i, Freq: make(map[int]int)}
}

func (g *Guard) Record(asleep, awake time.Time) {
	g.Sleep += awake.Sub(asleep)
	for i := asleep; i.Before(awake); i = i.Add(time.Minute) {
		g.Freq[i.Minute()] += 1
		if h, ok := g.Freq[g.Highest]; !ok || g.Freq[i.Minute()] > h {
			g.Highest = i.Minute()
		}
	}
}

type Time struct {
	Time time.Time
	Info string
}

func MakeTimeline(scanner *lib.AdventScanner) []*Time {
	var timeline []*Time
	pre := len(TimeLayout)
	for scanner.Scan() {
		text := scanner.Text()
		t, err := time.Parse(TimeLayout, text[:pre])
		if err != nil {
			log.Fatal(err)
		}
		timeline = append(timeline, &Time{Time: t, Info: text[pre+1:]})
	}
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].Time.Before(timeline[j].Time)
	})
	return timeline
}

func GetGuards(scanner *lib.AdventScanner) []*Guard {
	timeline := MakeTimeline(scanner)
	var guards []*Guard
	guardMap := make(map[string]*Guard)
	var asleep time.Time
	id := ""
	pre := len("Guard #")
	post := len(" begins shift")
	var guard *Guard
	for _, v := range timeline {
		switch v.Info {
		case "falls asleep":
			asleep = v.Time
		case "wakes up":
			guard.Record(asleep, v.Time)
		default:
			id = v.Info[pre:(len(v.Info) - post)]
			guard = guardMap[id]
			if guard == nil {
				guard = NewGuard(id)
				guardMap[id] = guard
				guards = append(guards, guard)
			}
		}
	}
	return guards
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := lib.NewAdventScanner(f)
	guards := GetGuards(scanner)
	var sleepy *Guard
	var consistent *Guard
	for _, v := range guards {
		if sleepy == nil || v.Sleep > sleepy.Sleep {
			sleepy = v
		}
		if consistent == nil || v.Freq[v.Highest] > consistent.Freq[consistent.Highest] {
			consistent = v
		}
	}
	fmt.Println("Most asleep:", int(sleepy.Id)*sleepy.Highest)
	fmt.Println("Most consistent:", int(consistent.Id)*consistent.Highest)
}
