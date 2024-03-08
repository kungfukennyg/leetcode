package linked

import (
	"testing"
)

var iterations = 10000
var marginOfError float64 = 5

func TestSolution(t *testing.T) {
	n := New(1)
	n.Append(2).Append(3)
	s := Constructor(n)
	counts := map[int]int{}
	for i := 0; i < iterations; i++ {
		v := s.GetRandom()
		counts[v] += 1
	}

	for k, v := range counts {
		percentile := (float64(v) / float64(iterations)) * 100
		expected := float64(100) / float64(s.len)
		if percentile+marginOfError < expected || percentile-marginOfError > expected {
			t.Fatalf("expected even distribution of all elements, got %v for %v", percentile, k)
		}
	}
}
