package file

import "os"

type DefaultFileWriter struct {}

// implements FileWriter interface Write()
func(fw *DefaultFileWriter) Write(data []byte) (int, error) {
	return os.File.W
}