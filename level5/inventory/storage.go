package inventory

import (
	"fmt"
	"io"
	"os"
)

func ReadInventory(path string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	fmt.Println("=== المخزون ===")
	return string(bytes), nil
}
