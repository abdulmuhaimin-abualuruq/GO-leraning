package catalog

import "testing"

func TestFindProductIndex(t *testing.T) {
	products := []Product{
		{Name: "Mouse", Price: 20, Quantity: 3},
		{Name: "Keyboard", Price: 15, Quantity: 2},
	}
	got := FindProductIndex(products, "Mouse")
	want := 0
	if got != want {
		t.Fatalf("FindProductIndex() = %d, want %d", got, want)
	}
	got = FindProductIndex(products, "Rust")
	want = -1
	if got != want {
		t.Fatalf("FindProductIndex() = %d, want %d", got, want)
	}
}

func TestSellByName(t *testing.T) {
	products := []Product{
		{Name: "Mouse", Price: 20, Quantity: 3},
		{Name: "Keyboard", Price: 15, Quantity: 2},
	}

	err := SellByName(products, "Mouse", 2)
	if err != nil {
		t.Fatalf("SellByName() = %v, want nil", err)
	}
	if products[0].Quantity != 1 {
		t.Fatalf("SellByName() = %d, want 1", products[0].Quantity)
	}

	err = SellByName(products, "Mouse", -1)
	if err == nil {
		t.Fatalf("SellByName() = nil, want error")
	}
	if products[0].Quantity != 1 {
		t.Fatalf("Quantity changed after failed SellByName: got %d, want 1", products[0].Quantity)
	}

	err = SellByName(products, "Mouse", 5)
	if err == nil {
		t.Fatalf("SellByName() = nil, want error")
	}
	if products[0].Quantity != 1 {
		t.Fatalf("Quantity changed after failed SellByName: got %d, want 1", products[0].Quantity)
	}

	err = SellByName(products, "Rust", 2)
	if err == nil {
		t.Fatalf("SellByName() = nil, want error")
	}
	if products[0].Quantity != 1 || products[1].Quantity != 2 {
		t.Fatalf("Stock corrupted after non-existent product sale attempt")
	}
}
