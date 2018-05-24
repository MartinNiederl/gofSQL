package gofSQL

import "os"

func ParseFile(path string) ([]Query, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return nil, nil
}
