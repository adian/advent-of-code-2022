package main

import (
	"adian.com/advent_of_code_2022/utils"
	"fmt"
	"math"
	"sort"
	"sync"
)

// This solution is overcomplicated because:
// * I was trying to learn goroutines.
// * I wanted to check how much performance I get If I use only pointers.

func main() {
	lines := utils.ReadFileLines("./day16/input.txt")
	fmt.Printf("Part 1 answer: %v\n", solvePart1(lines))
	fmt.Printf("Part 2 answer: %v\n", solvePart2(lines, 120_000))
}

func solvePart2(lines []string, maxGenerationSize int) int {
	startValve := parseInput(lines)

	snapshots := []*snapshot{
		{
			pressure:      0,
			opened:        make(map[string]*valve),
			currentValves: []*valve{&startValve, &startValve},
		},
	}
	for i := 0; i < 26; i++ {
		snapshots = filterBests(snapshots, maxGenerationSize)
		snapshots = nextMinute(snapshots)
		fmt.Printf("-> minute %d, number of snapshosts %d\n", i+1, len(snapshots))
	}

	return getBest(snapshots)
}

func solvePart1(lines []string) int {
	startValve := parseInput(lines)

	snapshots := []*snapshot{
		{
			pressure:      0,
			opened:        make(map[string]*valve),
			currentValves: []*valve{&startValve},
		},
	}
	for i := 0; i < 30; i++ {
		snapshots = filterBests(snapshots, 2_000)
		snapshots = nextMinute(snapshots)
		fmt.Printf("-> minute %d, number of snapshosts %d\n", i+1, len(snapshots))
	}

	return getBest(snapshots)
}

func getBest(snapshots []*snapshot) int {
	var highestResult int
	for _, s := range snapshots {
		currResult := s.pressure
		if currResult > highestResult {
			highestResult = currResult
		}
	}
	return highestResult
}

func filterBests(snapshots []*snapshot, maxSize int) []*snapshot {
	if len(snapshots) < maxSize {
		return snapshots
	}

	sort.Slice(snapshots, func(i, j int) bool {
		return snapshots[i].pressure > snapshots[j].pressure
	})

	snapshots = snapshots[:maxSize]
	return snapshots
}

func nextMinute(snapshots []*snapshot) []*snapshot {
	snapshots = increasePressure(snapshots)

	input := make(chan *snapshot, len(snapshots))
	defer close(input)
	for _, snap := range snapshots {
		input <- snap
	}

	poolSize := int(
		math.Min(float64(len(snapshots)), 75),
	)
	var wg sync.WaitGroup
	wg.Add(len(snapshots))

	output := make(chan []*snapshot, len(snapshots))
	for i := 0; i < poolSize; i++ {
		go func() {
			for s := range input {
				output <- processSnap(s)
				wg.Done()
			}
		}()
	}

	wg.Wait()
	close(output)

	var result []*snapshot
	for out := range output {
		result = append(result, out...)
	}

	return result
}

func processSnap(snap *snapshot) []*snapshot {
	var result []*snapshot

	for i := range snap.currentValves {
		if len(result) == 0 {
			temps := processSnapForCurrent(snap, i)
			result = append(result, temps...)
		} else {
			for _, s := range result {
				temps := processSnapForCurrent(s, i)
				result = append(result, temps...)
			}
		}
	}

	return result
}

func processSnapForCurrent(snap *snapshot, currIndex int) []*snapshot {
	var result []*snapshot
	curr := snap.currentValves[currIndex]

	if !snap.isOpened(curr) {
		temp := snap.copy()
		temp.open(curr)
		result = append(result, temp)
	}

	for _, nextValve := range curr.leadTo {
		temp := snap.copy()
		temp.currentValves[currIndex] = nextValve
		result = append(result, temp)
	}

	return result
}

func increasePressure(snapshots []*snapshot) []*snapshot {
	var temp []*snapshot
	for _, snap := range snapshots {
		for _, visited := range snap.opened {
			snap.pressure += visited.flowRate
		}
		temp = append(temp, snap)
	}

	return temp
}

type snapshot struct {
	pressure      int
	opened        map[string]*valve
	currentValves []*valve
}

func (snap *snapshot) isOpened(v *valve) bool {
	_, ok := snap.opened[v.name]
	return ok
}

func (snap *snapshot) open(v *valve) {
	snap.opened[v.name] = v
}

func (snap *snapshot) copy() *snapshot {
	newOpened := make(map[string]*valve, len(snap.opened))
	for k, v := range snap.opened {
		newOpened[k] = v
	}

	newCurrentValves := make([]*valve, len(snap.currentValves))
	copy(newCurrentValves, snap.currentValves)

	return &snapshot{
		pressure:      snap.pressure,
		opened:        newOpened,
		currentValves: newCurrentValves,
	}
}

type valve struct {
	name     string
	flowRate int
	leadTo   []*valve
}
