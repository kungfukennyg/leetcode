package insert

import "cmp"

type interval struct {
	left, right int
}

func (a interval) Compare(b interval) int {
	left, right := cmp.Compare(a.left, b.left), cmp.Compare(a.right, b.right)
	if left > right {
		return 1
	} else if left < right {
		return -1
	} else {
		return 0
	}
}

func (a interval) slice() []int {
	return []int{a.left, a.right}
}

func (a interval) merge(b interval) interval {
	return interval{left: min(a.left, b.left), right: max(a.left, a.right)}
}

func insert(ranges [][]int, newRange []int) [][]int {
	if len(ranges) == 0 || len(ranges[0]) != 2 || len(newRange) != 2 {
		return [][]int{}
	}

	newInterval := interval{left: newRange[0], right: newRange[1]}

	ret := [][]int{}
	for _, r := range intervals {
		ret = append(ret, r.slice())
	}
	return ret
}

func deduplicate(intervals []interval) []interval {
	ret := []interval{}
	for i, a := range intervals {
		for j, b := range intervals {
			if i == j {
				continue
			}
			if a != b {
				ret = append(ret, a)
			}
		}
	}
	return ret
}
