package inventory

import (
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
	return string(bytes), nil
}
