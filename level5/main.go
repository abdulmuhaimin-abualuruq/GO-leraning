package main

import (
	"example.com/academy/level5/inventory"
	"fmt"
)

func main() {
	content, err := inventory.ReadInventory("level5/data/products.txt")
	if err != nil {
		fmt.Println("Error reading inventory:", err)
		return
	}
	fmt.Println(content)
}
