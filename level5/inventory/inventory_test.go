package inventory

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseInventory(t *testing.T) {
	// 1. Success case with 3 Arabic products
	content := "KB-001 | لوحة مفاتيح | 12\nMS-002 | فأرة | 8\nMN-003 | شاشة | 4\n"
	products, err := ParseInventory(content)
	expected := []Product{
		{Code: "KB-001", Name: "لوحة مفاتيح", Quantity: 12},
		{Code: "MS-002", Name: "فأرة", Quantity: 8},
		{Code: "MN-003", Name: "شاشة", Quantity: 4},
	}
	if err != nil || !reflect.DeepEqual(products, expected) {
		t.Fatalf("ParseInventory failed or mismatch: err=%v, products=%+v", err, products)
	}

	// 2. Empty or whitespace content should return empty slice without error
	emptyInputs := []string{"", "   \n\n\t  \n"}
	for i := 0; i < len(emptyInputs); i++ {
		res, err := ParseInventory(emptyInputs[i])
		if err != nil || len(res) != 0 {
			t.Errorf("expected empty result without error for %q, got: %v, err=%v", emptyInputs[i], res, err)
		}
	}

	// 3. Reject invalid cases (non-numeric, negative, >20 runes, bad fields, empty fields)
	invalidInputs := []string{
		"KB-001 | لوحة مفاتيح | abc", // كمية غير رقمية
		"KB-001 | لوحة مفاتيح | -5",  // كمية سالبة
		"KB-001 | هذا الاسم العربي يتجاوز عشرين نقطة بالتأكيد | 10", // اسم أطول من 20 حرف
		"KB-001 | لوحة مفاتيح",              // حقول ناقصة
		"KB-001 | لوحة مفاتيح | 10 | extra", // حقول زائدة
		" | لوحة مفاتيح | 10",               // كود فارغ
		"KB-001 | | 10",                     // اسم فارغ
	}
	for i := 0; i < len(invalidInputs); i++ {
		if _, err := ParseInventory(invalidInputs[i]); err == nil {
			t.Errorf("expected error for invalid input: %q", invalidInputs[i])
		}
	}
}

func TestJSONStorage(t *testing.T) {
	products := []Product{
		{Code: "KB-001", Name: "لوحة مفاتيح", Quantity: 12},
		{Code: "MS-002", Name: "فأرة", Quantity: 8},
		{Code: "MN-003", Name: "شاشة", Quantity: 4},
	}
	path := filepath.Join(t.TempDir(), "inventory.json")

	// 1. Save and load products with matching Arabic data
	err := SaveJSON(path, products)
	if err != nil {
		t.Fatalf("SaveJSON failed: %v", err)
	}
	loaded, err := LoadJSON(path)
	if err != nil || !reflect.DeepEqual(loaded, products) {
		t.Fatalf("LoadJSON failed or data mismatch: err=%v, loaded=%+v", err, loaded)
	}

	// 2. Reject corrupt JSON file
	corruptPath := filepath.Join(t.TempDir(), "corrupt.json")
	if err := os.WriteFile(corruptPath, []byte(`{invalid`), 0644); err != nil {
		t.Fatalf("failed to prepare corrupt JSON: %v", err)
	}

	if _, err := LoadJSON(corruptPath); err == nil {
		t.Errorf("expected error for corrupt JSON")
	}

	// 3. Reject negative quantity without overwriting previously saved data
	invalid := []Product{{Code: "KB-001", Name: "لوحة مفاتيح", Quantity: -1}}
	if err := SaveJSON(path, invalid); err == nil {
		t.Errorf("expected error for negative quantity")
	}
	reloaded, err := LoadJSON(path)
	if err != nil || !reflect.DeepEqual(reloaded, products) {
		t.Fatalf("data was altered or unreadable after failed save: %v", err)
	}
}
