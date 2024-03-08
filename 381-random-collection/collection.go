package collection

import (
	"math/rand"
)

type RandomizedCollection struct {
	indices map[int][]int
	values  []int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		indices: map[int][]int{},
		values:  []int{},
	}
}

func (r *RandomizedCollection) Insert(v int) bool {
	idxs, ok := r.indices[v]
	r.values = append(r.values, v)
	idxs = append(idxs, len(r.values)-1)
	r.indices[v] = idxs
	return !ok
}

func (r *RandomizedCollection) Remove(v int) bool {
	idxs, ok := r.indices[v]
	if !ok {
		return false
	}

	// get the index to remove
	i := idxs[len(idxs)-1]
	lastIdx := len(r.values) - 1
	last := r.values[lastIdx]

	// swap element to remove with last element
	r.values[i] = last
	r.values = r.values[:lastIdx]

	// update index of swapped element
	lastIndices := r.indices[last]
	for j, idx := range lastIndices {
		if idx == lastIdx {
			lastIndices[j] = i
			break
		}
	}

	// remove the index from the map if it was the last occurrence
	if len(idxs) == 1 {
		delete(r.indices, v)
	} else {
		idxs[len(idxs)-1] = idxs[0]
		idxs = idxs[:len(idxs)-1]
		r.indices[v] = idxs
	}

	return true
}

func (r *RandomizedCollection) GetRandom() int {
	n := rand.Intn(len(r.values))
	return r.values[n]
}
