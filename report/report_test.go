package report

import (
	"testing"

	"example.com/academy/catalog"
)

func TestSum(t *testing.T) {
	items := []Valuable{
		catalog.Product{Price: 20, Quantity: 5},
		catalog.Service{Rate: 40, Sessions: 2},
		catalog.Product{Price: 15, Quantity: 2},
	}

	got := Sum(items)
	want := 210
	if got != want {
		t.Fatalf("Sum() = %d, want %d", got, want)
	}

	var empty []Valuable
	if got := Sum(empty); got != 0 {
		t.Fatalf("Sum(empty) = %d, want 0", got)
	}
}

func TestCountAbove(t *testing.T) {
	items := []Valuable{
		catalog.Product{Price: 20, Quantity: 5},
		catalog.Service{Rate: 40, Sessions: 2},
		catalog.Product{Price: 15, Quantity: 2},
	}

	if got := CountAbove(items, 50); got != 2 {
		t.Fatalf("CountAbove(items, 50) = %d, want 2", got)
	}

	if got := CountAbove(items, 100); got != 0 {
		t.Fatalf("CountAbove(items, 100) = %d, want 0", got)
	}

	var empty []Valuable
	if got := CountAbove(empty, 50); got != 0 {
		t.Fatalf("CountAbove(empty, 50) = %d, want 0", got)
	}
}
