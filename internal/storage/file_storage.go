package storage

import (
	"errors"
	"os"
)

type FileStorage struct{}

func (*FileStorage) SaveLocally(fileName string, data []byte) (int, error) {
	file, fileErr := os.Create(fileName)

	if fileErr != nil {
		defer file.Close()
		return 0, errors.New("Error creating " + fileName)
	}

	defer file.Close() // Close the file automatically once I/O operations are finished

	size, writeErr := file.Write(data)

	if writeErr != nil {
		return 0, errors.New("error writing data locally")
	}

	return size, nil
}
