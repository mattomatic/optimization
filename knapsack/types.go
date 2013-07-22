package main

import (
	"bytes"
	"fmt"
)

type Item struct {
	Id     int
	Value  int
	Weight int
}

type Problem struct {
	Size     int
	Capacity int
	Items    []Item
}

type Solution struct {
	Objective int
	Optimal   bool
	Items     ItemSet
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

	for _, item := range s.Problem.Items {
		_, ok := s.Items[item]

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

	for item := range s.Items {
		fmt.Println("\t", item)
	}

	fmt.Println("---------")
	fmt.Println("output:")
	fmt.Println(s.String())
}
