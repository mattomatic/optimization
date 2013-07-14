package main

import (
	"sort"
)

// allow items to be sorted
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

func greedySolve(p *Problem) *Solution {
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
