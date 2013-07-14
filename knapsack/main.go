package main

import (
	"flag"
	"fmt"
)

func init() {
	flag.Parse()
}

func main() {
	if flag.NArg() != 1 {
		fmt.Println("usage: ./knapsack <inputFile>")
	}

	filename := flag.Arg(0)
	p := LoadProblemFile(filename)
	solution := greedySolve(p)
	fmt.Println(solution.String())
}
