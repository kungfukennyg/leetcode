package merge

import (
	"reflect"
	"testing"
)

func Test_merge2(t *testing.T) {
	type args struct {
		nums1, nums2 []int
		m, n         int
	}

	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "One",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				nums2: []int{2, 5, 6},
				m:     3,
				n:     3,
			},
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Two",
			args: args{
				nums1: []int{1},
				nums2: []int{},
				m:     1,
				n:     1,
			},
			expected: []int{1},
		},
		{
			name: "Three",
			args: args{
				nums1: []int{0},
				nums2: []int{1},
				m:     0,
				n:     1,
			},
			expected: []int{1},
		},
		{
			name: "Four",
			args: args{
				nums1: []int{4, 5, 6, 0, 0, 0},
				nums2: []int{1, 2, 3},
				m:     3,
				n:     3,
			},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Negative",
			args: args{
				nums1: []int{-1, 0, 0, 3, 3, 3, 0, 0, 0},
				nums2: []int{1, 2, 2},
				m:     6,
				n:     3,
			},
			expected: []int{-1, 0, 0, 1, 2, 2, 3, 3, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			merge(tc.args.nums1, tc.args.m, tc.args.nums2, tc.args.n)
			if !reflect.DeepEqual(tc.args.nums1, tc.expected) {
				t.Errorf("%s: expected: %v, actual: %v\n", tc.name, tc.expected, tc.args.nums1)
			}
		})
	}
}
