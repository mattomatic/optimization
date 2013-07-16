package main

import (
    "https://github.com/hotei/bits"
)

type Real float64

type node struct {
    explored bits.BitField
    status bits.BitField
    estimate Real
    cost int
}

// Solve the problem using dynamic programming by building a matrix
// of optimal solutions from the ground up. 
func dynamicSolve(p *Problem) *Solution {
    

	solution := &Solution{}
	return solution
}
