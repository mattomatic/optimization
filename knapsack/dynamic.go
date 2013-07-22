package main

type node struct {
	included ItemSet
	excluded ItemSet
	estimate int
	weight   int
	value  int

}

func (n *node) getSolution(p *Problem) *Solution {
    return &Solution{n.value, n.value == n.estimate, n.included, *p}
}

func newNode(problem *Problem, included, excluded ItemSet) *node {
	return &node{included, excluded, getEstimate(problem, included, excluded), getWeight(included), getValue(included)}
}

// Solve the problem using dynamic programming by building a matrix
// of optimal solutions from the ground up.
func dynamicSolve(p *Problem) *Solution {
	sortItems(p.Items)

	included := NewItemSet()
	excluded := NewItemSet(p.Items...)
	best := newNode(p, included, excluded)
	backTrack(p, best, best)
	return best.getSolution(p)
}

func backTrack(p *Problem, n *node, best *node) {
	// If we're carrying too much, get out!
	if n.weight > p.Capacity {
		return
	}
	
	// If our best scenario is worse than an existing thing, get out!
	if n.estimate < best.value {
	    return
	}
	
	// If our value is better than the previous best, replace it
	if n.value > best.value {
	    *best = *n
	}

	for item := range n.excluded {
		backTrack(p, newNode(p, n.included.add(item), n.excluded.remove(item)), best)
	}
}

// Add up the weight of the items in the ItemSet
func getWeight(set ItemSet) int {
	weight := 0

	for item := range set {
		weight += item.Weight
	}

	return weight
}

// Add up the value of the items in the ItemSet
func getValue(set ItemSet) int {
	value := 0

	for item := range set {
		value += item.Value
	}

	return value
}

// Add up the value of all items we may possibly include
// in the future. This estimate is completely optimistic
// as we are allowed to include even partial items.
func getEstimate(p *Problem, included, excluded ItemSet) int {
	capacity := p.Capacity
	estimate := 0

	for _, item := range p.Items {
		_, okIncluded := included[item]
		_, okExcluded := excluded[item]

		if okIncluded || okExcluded {
			delta := getMin(capacity, item.Weight)
			capacity -= delta
			estimate += delta * item.Value
		}
	}

	return estimate
}

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
