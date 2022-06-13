package utils

import (
	"os"
)

func ReadFile(path string) []byte {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return dat
}

func WriteFile(path string, data []byte) {
	os.WriteFile(path, data, os.FileMode.Perm(0o644))
}
