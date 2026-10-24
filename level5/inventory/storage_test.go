package inventory

import (
	"testing"
)

func TestReadInventorySuccess(t *testing.T) {
	read, err := ReadInventory("../data/products.txt")
	if err != nil {
		t.Fatalf("❌❌ Error reading inventory: %v", err)
	}
	expected := "KB-001 | لوحة مفاتيح | 12\nMS-002 | فأرة | 8\nMN-003 | شاشة | 4\n"
	if read != expected {
		t.Fatalf("❌❌ Expected content:\n%q\nGot:\n%q", expected, read)
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
