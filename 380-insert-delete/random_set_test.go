package random_set

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	rs := Constructor()
	r := rs.Insert(1)
	if !r {
		t.Fatal()
	}
	r = rs.Remove(2)
	if r {
		t.Fatal()
	}
	r = rs.Insert(2)
	if !r {
		t.Fatal()
	}
	n := rs.GetRandom()
	if n != 1 && n != 2 {
		t.Fatal()
	}
	r = rs.Remove(1)
	if !r {
		t.Fatal()
	}
	r = rs.Insert(2)
	if r {
		t.Fatal()
	}

	for i := 0; i < 5; i++ {
		n = rs.GetRandom()
		if n != 2 {
			t.Fatalf("expected GetRandom to always return 2, got %d", n)
		}
	}
}

func TestRandomizedSetZero(t *testing.T) {
	rs := Constructor()
	rs.Insert(0)
	rs.Insert(1)
	rs.Remove(0)
	rs.Insert(2)
	rs.Remove(1)

	if r := rs.GetRandom(); r != 2 {
		t.Fatalf("expected GetRandom to return 2, got %d", r)
	}
}

func TestRandomizedSetNegatives(t *testing.T) {
	rs := Constructor()
	rs.Insert(3)
	rs.Insert(-2)
	rs.Remove(2)
	rs.Insert(1)
	rs.Insert(-3)
	rs.Insert(-2)
	rs.Remove(-2)
	rs.Remove(3)
	rs.Insert(-1)
	rs.Remove(-3)
	rs.Insert(1)
	rs.Insert(-2)
	rs.Insert(-2)
	rs.Insert(-2)
	rs.Insert(1)
	r := rs.GetRandom()

	if r != -1 && r != -2 {
		t.Fatalf("expected GetRandom to return 2, got %d", r)
	}
}

func BenchmarkRandomizedSet_Insert(b *testing.B) {
	for _, n := range inputSizes {
		rs := setup(b, n)
		b.Run(fmt.Sprintf("Insert-%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs.Insert(i)
			}
		})
	}
}
func BenchmarkRandomizedSet_Remove(b *testing.B) {
	for _, n := range inputSizes {
		rs := setup(b, n)
		b.Run(fmt.Sprintf("Remove-%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs.Remove(i)
			}
		})
	}
}
func BenchmarkRandomizedSet_GetRandom(b *testing.B) {
	for _, n := range inputSizes {
		rs := setup(b, n)
		b.Run(fmt.Sprintf("GetRandom-%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				rs.GetRandom()
			}
		})
	}
}

var inputSizes = []int{10, 100, 1000, 10000, 100_000, 1_000_000}

func setup(b *testing.B, n int) *RandomizedSet {
	b.Helper()
	rs := Constructor()
	for i := 0; i < n; i++ {
		rs.Insert(rand.Int())
	}
	return &rs
}
