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
