package main

type ItemSet map[Item]bool

func NewItemSet(items ...Item) ItemSet {
	set := make(map[Item]bool)

	for _, item := range items {
		set[item] = true
	}

	return set
}

func (s ItemSet) copy() ItemSet {
	set := make(map[Item]bool)

	for k, v := range s {
		set[k] = v
	}

	return set
}

func (s ItemSet) add(item Item) ItemSet {
	set := s.copy()
	set[item] = true
	return set
}

func (s ItemSet) remove(item Item) ItemSet {
	set := s.copy()
	delete(set, item)
	return set
}
