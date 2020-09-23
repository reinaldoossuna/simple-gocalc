package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push(1.0)
	s.Push(2.0)
	want := 3.0
	s.Push(want)

	if got := s.Pop(); got.Value != want {
		t.Errorf("Stack Pop = %v, want %v", got, want)
	}
}

func TestShuffle(t *testing.T) {
	s := Stack{}
	s.Push(1.0)
	want := 2.0
	s.Push(want)
	s.Push(3.0)
	s.Shuffle()

	if got := s.Pop(); got.Value != want {
		t.Errorf("Stack Pop = %v, want %v", got, want)
	}
}

func TestPick(t *testing.T) {
	s := Stack{}
	want := 1.0
	s.Push(want)
	s.Push(2.0)
	s.Push(3.0)

	if got := s.Pick(2); got.Value != want {
		t.Errorf("Stack Pop = %v, want %v", got, want)
	}
}

func TestPreview(t *testing.T) {
	s := Stack{}
	n := 4
	for i := 0; i < n; i++ {
		s.Push(float64(i))
	}
	got := s.Preview(4)
	for i := 0; i < n; i++ {
		t.Errorf("Preview = %v, want %v", got[i], float64(n - i - 1))
	}
}
