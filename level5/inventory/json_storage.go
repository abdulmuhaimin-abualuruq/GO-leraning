package inventory

import (
	"encoding/json"
	"os"
)

// SaveJSON validates all products, encodes them to formatted JSON, and writes to path.
func SaveJSON(path string, products []Product) error {

	if err := ValidateInventory(products); err != nil {
		return err
	}

	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// LoadJSON reads the JSON file at path, deserializes it into []Product, and validates the data.
func LoadJSON(path string) ([]Product, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	products := []Product{}
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	if err := ValidateInventory(products); err != nil {
		return nil, err
	}

	return products, nil
}
