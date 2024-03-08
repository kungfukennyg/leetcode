package shuffle

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		ops []string
		in  []int
	}

	tcs := []testCase{
		{
			ops: []string{"shuffle", "reset", "shuffle"},
			in:  []int{1, 2, 3},
		},
		{
			ops: []string{"shuffle"},
			in:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			s := Constructor(tc.in)
		OUTER:
			for _, op := range tc.ops {
				got := exec(op, &s)
				if op == "reset" {
					if !reflect.DeepEqual(got, tc.in) {
						t.Fatalf("expected value to be %v, got %v", tc.in, got)
					}
					continue
				}
				for _, n := range tc.in {
					for _, m := range got {
						if n == m {
							continue OUTER
						}
					}
					t.Fatalf("expected value to contain all of input values %v, got: %v", tc.in, got)
				}
			}
		})
	}
}

func exec(op string, s *Solution) []int {
	switch op {
	case "shuffle":
		return s.Shuffle()
	case "reset":
		return s.Reset()
	default:
		panic(fmt.Sprintf("unrecognized op %s", op))
	}
}
