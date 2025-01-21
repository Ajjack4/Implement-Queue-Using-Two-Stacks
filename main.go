package main

import (
	"fmt"
)

type Queue struct {
	s1 []int
	s2 []int
}

func (q *Queue) Add(v int) {

	for len(q.s2) > 0 {
		q.s1 = append(q.s1, q.s2[len(q.s2)-1])
		q.s2 = q.s2[:len(q.s2)-1]
	}
	q.s1 = append(q.s1, v)
	for len(q.s1) > 0 {
		q.s2 = append(q.s2, q.s1[len(q.s1)-1])
		q.s1 = q.s1[:len(q.s1)-1]
	}
}
func (q *Queue) Pop() int {
	v := q.s2[len(q.s2)-1]
	q.s2 = q.s2[:len(q.s2)-1]
	return v
}
func main() {
	q := Queue{
		s1: []int{},
		s2: []int{1, 2},
	}
	q.Add(3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q)
}
