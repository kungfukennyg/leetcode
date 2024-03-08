package linked

import "math/rand"

type Solution struct {
	len  int
	Head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) Append(v int) *ListNode {
	n.Next = &ListNode{
		Val: v,
	}
	return n.Next
}

func New(v int) *ListNode {
	return &ListNode{Val: v}
}

func Constructor(h *ListNode) Solution {
	s := Solution{Head: h, len: 0}
	n := h
	for {
		if n == nil {
			break
		}
		n = n.Next
		s.len++
	}
	return s
}

func (s *Solution) GetRandom() int {
	if s.len == 0 {
		return 0
	}

	r := rand.Intn(s.len)
	n := s.Head
	i := 0
	for {
		if i == r {
			return n.Val
		}
		n = n.Next
		i++
	}
}
