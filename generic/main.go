package main

import (
	"fmt"
	"sort"
)

type Node[T any] struct {
	left  *Node[T]
	right *Node[T]
	val   T
}

type Tree[T any] struct {
	compare func(T, T) int
	root    *Node[T]
}

func (t *Tree[T]) find(val T) *Node[T] {
	iter := t.root

	for iter != nil {
		switch cmp := t.compare(val, iter.val); {
		case cmp < 0:
			iter = iter.left
		case cmp > 0:
			iter = iter.right
		default:
			return iter
		}
	}
	return iter
}

type MySlice[T any] struct {
	s    []T
	less func(T, T) bool
}

func (s MySlice[T]) Len() int {
	return len(s.s)
}

func (s MySlice[T]) Less(i, j int) bool {
	return s.less(s.s[i], s.s[j])
}

func (s MySlice[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
	res := make([]Key, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

func less(x, y int) bool {
	return x < y
}

func compare(x, y int) int {
	if x > y {
		return 1
	} else if x < y {
		return -1
	}
	return 0
}

func main() {
	m1 := map[int]int{
		1: 1,
		2: 2,
	}
	m2 := map[string]string{
		"A": "A",
		"B": "B",
	}
	fmt.Println(MapKeys(m1), MapKeys(m2))

	t := Tree[int]{
		compare: compare,
		root:    &Node[int]{val: 5},
	}
	fmt.Println(t.find(5))

	ms := &MySlice[int]{
		s:    []int{3, 2, 1},
		less: less,
	}
	sort.Sort(ms)
	fmt.Println(*ms)
}
