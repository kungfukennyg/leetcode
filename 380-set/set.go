package set

import (
	"math/rand"
)

// RandomizedSet is a map-backed set providing O(1) insertion, deletion, and random lookup time
// at the cost of a larger memory footprint
type RandomizedSet struct {
	// the set itself
	s map[int]struct{}
	// a slice of keys providing constant time random access
	keys []int
	// a map of values to index in keys slice, allowing constant time deletions
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		s:       map[int]struct{}{},
		keys:    []int{},
		indices: map[int]int{},
	}
}

func (rs *RandomizedSet) Insert(v int) bool {
	_, ok := rs.s[v]
	rs.s[v] = struct{}{}
	if !ok {
		idx := len(rs.keys)
		rs.keys = append(rs.keys, v)
		rs.indices[v] = idx
	}
	return !ok
}

func (rs *RandomizedSet) Remove(v int) bool {
	_, ok := rs.s[v]
	delete(rs.s, v)
	if ok {
		idx := rs.indices[v]
		delete(rs.indices, v)
		// replace this key with last element and trim end of slice
		// for O(1) removal
		lastKey := rs.keys[len(rs.keys)-1]
		rs.keys[idx] = lastKey
		rs.keys = rs.keys[:len(rs.keys)-1]
		rs.indices[lastKey] = idx
	}
	return ok
}

func (rs *RandomizedSet) GetRandom() int {
	rand := rand.Intn(len(rs.s))
	return rs.keys[rand]
}
