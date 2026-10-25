package inventory

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Product represents an inventory item.
type Product struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// ValidateProduct checks product rules using switch and ||
func ValidateProduct(p Product) error {
	runes := utf8.RuneCountInString(p.Name)
	switch {
	case strings.TrimSpace(p.Code) == "" || strings.TrimSpace(p.Name) == "":
		return fmt.Errorf("كود المنتج واسمه مطلوبان")
	case runes < 1 || runes > 20:
		return fmt.Errorf("طول اسم المنتج غير صالح: يجب أن يكون بين 1 و 20 نقطة Unicode")
	case p.Quantity < 0:
		return fmt.Errorf("الكمية لا يمكن أن تكون سالبة")
	default:
		return nil
	}
}

// ParseInventory parses text content into a slice of Product.
func ParseInventory(content string) ([]Product, error) {
	products := []Product{}
	lines := strings.Split(content, "\n")

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		fields := strings.Split(line, "|")
		if len(fields) != 3 {
			return nil, fmt.Errorf("السطر %d غير صالح: يجب أن يحتوي على 3 حقول", i+1)
		}

		qty, err := strconv.Atoi(strings.TrimSpace(fields[2]))
		if err != nil {
			return nil, fmt.Errorf("السطر %d: الكمية غير رقمية", i+1)
		}

		p := Product{
			Code:     strings.TrimSpace(fields[0]),
			Name:     strings.TrimSpace(fields[1]),
			Quantity: qty,
		}

		if err := ValidateProduct(p); err != nil {
			return nil, fmt.Errorf("السطر %d: %w", i+1, err)
		}

		products = append(products, p)
	}

	if err := ValidateInventory(products); err != nil {
		return nil, err
	}

	return products, nil
}

func ValidateInventory(products []Product) error {
	seen := map[string]bool{}

	for i := 0; i < len(products); i++ {
		p := products[i]

		if err := ValidateProduct(p); err != nil {
			return err
		}
		code := strings.TrimSpace(p.Code)

		if seen[code] {
			return fmt.Errorf("كود المنتج %s مكرر", code)
		}

		seen[code] = true
	}

	return nil
}
