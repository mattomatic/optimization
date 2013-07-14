package main

import (
	"bytes"
	"fmt"
	"sort"
)

type Solution struct {
	Objective int
	Optimal   bool
	Items     map[int]bool
	Problem
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (s *Solution) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("%d %d\n", s.Objective, btoi(s.Optimal)))

	for id, _ := range s.Problem.Items {
		_, ok := s.Items[id]

		if ok {
			buffer.WriteString("1 ")
		} else {
			buffer.WriteString("0 ")
		}
	}

	buffer.WriteString("\n")

	return buffer.String()
}

func (s *Solution) Print() {
	fmt.Printf("solution objective(%d) optimal(%t)\n", s.Objective, s.Optimal)

	for id, _ := range s.Items {
		fmt.Println("\t", s.Problem.Items[id])
	}

	fmt.Println("---------")
	fmt.Println("output:")
	fmt.Println(s.String())
}

// allow items to be sortable
type itemSorter struct {
	items []Item
}

func (s *itemSorter) Swap(a, b int) {
	tmp := s.items[a]
	s.items[a] = s.items[b]
	s.items[b] = tmp
}

func (s *itemSorter) Less(a, b int) bool {
	aval := float64(s.items[a].Value) / float64(s.items[a].Weight)
	bval := float64(s.items[b].Value) / float64(s.items[b].Weight)
	return aval < bval
}

func (s *itemSorter) Len() int {
	return len(s.items)
}

func solve(p *Problem) *Solution {
	sorter := &itemSorter{p.Items}
	sort.Sort(sort.Reverse(sorter))

	capacity := p.Capacity
	solution := &Solution{Problem: *p, Items: make(map[int]bool)}

	for _, item := range p.Items {
		if item.Weight < capacity {
			solution.Items[item.Id] = true
			solution.Objective += item.Value
			capacity -= item.Weight
		}
	}

	if capacity == 0 {
		solution.Optimal = true
	}

	return solution
}
