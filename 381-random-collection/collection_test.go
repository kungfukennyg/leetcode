package collection

import (
	"reflect"
	"testing"
)

func TestCollection(t *testing.T) {
	tcs := []testCase{
		{
			name:     "1",
			ops:      []string{"insert", "insert", "insert", "getRandom", "remove", "getRandom"},
			args:     []int{1, 1, 2, 0, 1, 0},
			expected: []any{true, false, true, nil, true, nil},
		},
		{
			name:     "2",
			ops:      []string{"insert", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "remove"},
			args:     []int{4, 3, 4, 2, 4, 4, 3, 4, 4},
			expected: []any{true, true, false, true, false, true, true, true, true},
		},
		{
			name:     "3",
			ops:      []string{"insert", "insert", "insert", "insert", "insert", "insert", "remove", "remove", "remove", "remove", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"},
			args:     []int{10, 10, 20, 20, 30, 30, 10, 10, 30, 30, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expected: []any{true, false, true, false, true, false, true, true, true, true, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			rs := Constructor()
			elementCounts := map[int]int{}
			for i, op := range tc.ops {
				r := exec(t, &rs, elementCounts, op, tc.args[i])
				expected := tc.expected[i]
				if expected == nil && op == "getRandom" {
					for n, count := range elementCounts {
						if r == n && count <= 0 {
							t.Fatalf("got unexpected value %v", r)
						}
					}
				} else if !reflect.DeepEqual(r, expected) {
					t.Fatalf("[%d] expected %v, got %v", i, expected, r)
				}
			}
		})
	}
}

func exec(t *testing.T, rs *RandomizedCollection, elementCounts map[int]int, op string, arg int) any {
	t.Helper()
	switch op {
	case "insert":
		elementCounts[arg] += 1
		return rs.Insert(arg)
	case "remove":
		elementCounts[arg] -= 1
		return rs.Remove(arg)
	case "getRandom":
		return rs.GetRandom()
	default:
		t.Fatalf("unexpected call op %s", op)
	}

	return nil
}

type testCase struct {
	name     string
	ops      []string
	args     []int
	expected []any
}
