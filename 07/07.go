package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in, out, err := getInput("07.input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(SortSteps(in, out))
	fmt.Println(SortStepsWithTime(in, out, 5, 60))
}

func getInput(filename string) (map[string][]string, map[string][]string, error) {
	in := make(map[string][]string)
	out := make(map[string][]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var from, to string
		fmt.Sscanf(scanner.Text(), "Step %s must be finished before step %s can begin.", &from, &to)
		in[to] = append(in[to], from)
		out[from] = append(out[from], to)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return in, out, nil
}

// SortSteps takes a map of in-edges and a map of out-edges, and returns a
// topological sort with tie-breakers by ascending alphabetical order.
func SortSteps(in map[string][]string, out map[string][]string) string {
	visited := make(map[string]bool)
	sequence := []string{}

	// start from steps that have outbound edges and no inbound edges
	toVisit := []string{}
	for step := range out {
		if len(in[step]) == 0 {
			toVisit = append(toVisit, step)
		}
	}
	// this is a priority queue, which must only contain the available
	// steps (all pre-requisites are fulfilled)
	sort.Strings(toVisit)

	for len(toVisit) > 0 {
		// grab the first available step in alphabetical order
		now := toVisit[0]
		toVisit = toVisit[1:]

		// if it's already visited for some reason, just ignore it
		if visited[now] {
			continue
		}
		visited[now] = true
		sequence = append(sequence, now)

		// check if next steps are ready to be visited
		for _, next := range out[now] {
			ready := true
			for _, i := range in[next] {
				if !visited[i] {
					ready = false
					break
				}
			}
			if ready {
				toVisit = append(toVisit, next)
				sort.Strings(toVisit)
			}
		}
	}

	return strings.Join(sequence, "")
}

// NextStep encapsulates a next step that's available at a certain time onwards
// (inclusive).
type NextStep struct {
	step string
	time int // eligible time
}

// NextSteps for use as sort.Interface
type NextSteps []NextStep

// sort.Interface
func (members NextSteps) Len() int {
	return len(members)
}
func (members NextSteps) Swap(i, j int) {
	members[i], members[j] = members[j], members[i]
}
func (members NextSteps) Less(i, j int) bool {
	if members[i].time < members[j].time {
		return true
	}
	if members[i].time > members[j].time {
		return false
	}
	return strings.Compare(members[i].step, members[j].step) == -1
}

// SortStepsWithTime takes a map of in-edges and a map of out-edges, and returns
// a topological sort with tie-breakers by ascending alphabetical order, and
// factoring in time taken for each step.
func SortStepsWithTime(in map[string][]string, out map[string][]string, workerNum int, stepDelay int) int {
	// additional tracking required for
	// (i) when tasks are completed (affects next task availability)
	// (ii) when worker is available for next task

	// visited is int of time that step is completed: the nil value 0 is
	// impossible since all steps take at least 61 seconds.
	visited := make(map[string]int)
	// second after which the worker will be available
	workerReady := map[int]int{}

	// instantiate the workers, as otherwise they don't exist
	for i := 0; i < workerNum; i++ {
		workerReady[i] = 0
	}

	// start from steps that have outbound edges and no inbound edges
	toVisit := []NextStep{}
	for step := range out {
		if len(in[step]) == 0 {
			toVisit = append(toVisit, NextStep{step: step, time: 0})
		}
	}
	// cast to NextSteps type to use custom sort.Interface
	sort.Sort(NextSteps(toVisit))
	fmt.Printf("At start: %v\n", toVisit)

	for len(toVisit) > 0 {
		now := toVisit[0]
		toVisit = toVisit[1:]

		if visited[now.step] != 0 {
			fmt.Printf("[%s] visiting, but already visited, skip\n", now.step)
			continue
		}
		fmt.Printf("[%s] visiting\n", now.step)

		// locate next available worker
		nextWorkerStartTime := -1
		nextWorkerID := -1
		for id, workerStartTime := range workerReady {
			if nextWorkerStartTime == -1 || workerStartTime < nextWorkerStartTime {
				nextWorkerStartTime = workerStartTime
				nextWorkerID = id
			}
		}
		fmt.Printf("[%s] next available worker %d at time %d\n", now.step, nextWorkerID, nextWorkerStartTime)

		// compute when this step will be completed, update worker and step.
		stepDuration := stepDelay + 1 + int(now.step[0]-'A')
		if nextWorkerStartTime < now.time {
			nextWorkerStartTime = now.time
		}
		nextWorkerEndTime := nextWorkerStartTime + stepDuration
		workerReady[nextWorkerID] = nextWorkerEndTime
		visited[now.step] = nextWorkerEndTime
		fmt.Printf("[%s] this step will start at %d and end at %d\n", now.step, nextWorkerStartTime, nextWorkerEndTime)

		// consider next steps: they are available only when the ending time
		// of all previous steps are known, since only then can the earliest
		// start time be found
		for _, next := range out[now.step] {
			ready := true
			readyTime := 0
			for _, i := range in[next] {
				if visited[i] == 0 {
					ready = false
				} else {
					if visited[i] > readyTime {
						readyTime = visited[i]
					}
				}
			}
			if ready {
				toVisit = append(toVisit, NextStep{step: next, time: readyTime})

			}
		}
		sort.Sort(NextSteps(toVisit))
	}

	maxTime := 0
	for _, time := range visited {
		if time > maxTime {
			maxTime = time
		}
	}

	return maxTime
}
