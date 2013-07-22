package main

import (
	"flag"
	"fmt"

)

func main() {
	solver := flag.String("solver", "dynamic", "specify solver (greedy|dynamic)")
	flag.Parse()
	filename := flag.Arg(0)
	p := LoadProblemFile(filename)

	var solution *Solution
	
	StartCPUProfile("prof")

	switch *solver {
	case "greedy":
		solution = greedySolve(p)
	case "dynamic":
		solution = dynamicSolve(p)
	default:
		panic("unrecognized solver")
	}
	
	StopCPUProfile()

	fmt.Println(solution.String())
}
