package catalog

import "errors"

func FindProductIndex(products []Product, name string) int {
	for i := range products {
		if products[i].Name == name {
			return i
		}
	}
	return -1
}
func SellByName(products []Product, name string, amount int) error {
	idx := FindProductIndex(products, name)
	if idx == -1 {
		return errors.New("product not found")
	}
	return products[idx].Sell(amount)
}
