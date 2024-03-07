package random_set

import (
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
