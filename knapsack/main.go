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
	solution := solve(p)
	solution.Print()
}
