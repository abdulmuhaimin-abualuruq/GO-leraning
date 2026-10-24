package inventory

import (
	"testing"
)

func TestReadInventorySuccess(t *testing.T) {
	read, err := ReadInventory("../data/products.txt")
	if err != nil {
		t.Fatalf("❌❌ Error reading inventory: %v", err)
	}
	if read == "" {
		t.Fatal("❌❌ Expected content but got empty string")
	}
}

func TestReadInventoryMissing(t *testing.T) {
	read, err := ReadInventory("../data/missing.txt")
	if err == nil {
		t.Fatal("❌❌ Expected error for missing file, but got nil")
	}
	if read != "" {
		t.Fatalf("❌❌ Expected empty content on error, got: %v", read)
	}
}
