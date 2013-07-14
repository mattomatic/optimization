package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadProblemFile(filename string) *Problem {
	fp, err := os.Open(filename)

	if err != nil {
		panic("could not open file")
	}

	reader := bufio.NewReader(fp)
	return loadProblem(reader)
}

func LoadProblem() *Problem {
	reader := bufio.NewReader(os.Stdin)
	return loadProblem(reader)
}

func loadProblem(reader *bufio.Reader) *Problem {
	lhs := 0
	rhs := 0
	count, err := fmt.Fscan(reader, &lhs, &rhs)

	if count != 2 || err != nil {
		panic("not enough items in header line")
	}

	p := &Problem{lhs, rhs, make([]Item, 0, 0)}

	for count, err := fmt.Fscan(reader, &lhs, &rhs); count == 2 && err == nil; count, err = fmt.Fscan(reader, &lhs, &rhs) {
		id := len(p.Items)
		p.Items = append(p.Items, Item{id, lhs, rhs})
	}

	return p
}
