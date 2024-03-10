package parser

import "strconv"

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	v        int
	children []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return len(n.children) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	if !n.IsInteger() {
		panic("GetInteger call on non integer object")
	}
	return n.v
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(v int) {
	n.v = v
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(e NestedInteger) {
	n.children = append(n.children, &e)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.children
}

func deserialize(s string) *NestedInteger {
	root := NestedInteger{}

	var negative bool
	var num []byte
	cur := &root
	var prev *NestedInteger
	var multi bool
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isNumber(c) {
			// hit some part of a number
			num = append([]byte(num), c)
			continue
		}

		switch string(c) {
		case "-":
			negative = true
		case "[":
			cur = &NestedInteger{}
		case "]":
			multi = false
			// completed a sublist
		case ",":
			if len(num) != 0 {
				// completed a number
				n := parse(string(num), negative)
				if multi {
					cur.Add(NestedInteger{v: n})
				} else {
					cur.SetInteger(n)
				}
			}
		}
	}

	if len(num) != 0 {
		// make sure we didn't miss a straggler
		n := parse(string(num), negative)
		if multi {
			cur.Add(NestedInteger{v: n})
		} else {
			cur.SetInteger(n)
		}
	}

	return &root
}

func isNumber(c byte) bool {
	_, err := strconv.ParseInt(string(c), 10, 64)
	return err == nil
}

func parse(num string, negative bool) int {
	n, _ := strconv.ParseInt(num, 10, 64)
	if negative {
		n = -n
	}
	return int(n)
}
