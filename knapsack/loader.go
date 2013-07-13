package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
	Value  int
	Weight int
}

type Problem struct {
	Size     int
	Capacity int
	Items    []Item
}

func LoadProblem(filename string) *Problem {
    fp, err := os.Open(filename)
    
    if err != nil {
        panic("could not open file")
    }
    
    reader := bufio.NewReader(fp)
    
    lhs := 0
    rhs := 0
    count, err := fmt.Fscan(reader, &lhs, &rhs)
    
    if count != 2 || err != nil {
        panic("not enough items in header line")
    }
    
    p := &Problem{lhs, rhs, make([]Item, 0, 0)}
    
    for count, err := fmt.Fscan(reader, &lhs, &rhs); count == 2 && err == nil; count, err = fmt.Fscan(reader, &lhs, &rhs) {
        p.Items = append(p.Items, Item{lhs, rhs})
    }
    
    return p
}
    