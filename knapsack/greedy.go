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

// sort the items by f(x) = x.Value / x.Weight
func sortItems(items []Item) {
	sorter := &itemSorter{items}
	sort.Sort(sort.Reverse(sorter))
}

// The greedy solver sorts the items by f(x) = x.Value / x.Weight
// and chooses as many items as will fit into the knapsack. Note
// that it will not necessarily choose the first N items, as
// some items may not fit even though f(x_i) is higher than f(x_i+1)
func greedySolve(p *Problem) *Solution {
	sortItems(p.Items)

	capacity := p.Capacity
	solution := &Solution{Problem: *p, Items: NewItemSet()}

	for _, item := range p.Items {
		if item.Weight < capacity {
			solution.Items.add(item)
			solution.Objective += item.Value
			capacity -= item.Weight
		}
	}

	if capacity == 0 {
		solution.Optimal = true
	}

	return solution
}
