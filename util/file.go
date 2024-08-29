package util

import (
	"os"
)

// FetchFileContent returns the byte array of a file located at the specified path or in the parent directory
func FetchFileContent(filePath string) (fileContent []byte) {
	fileContent, readError := os.ReadFile(filePath)
	if readError != nil {
		fileContent, readError = os.ReadFile("../" + filePath)
	}

	if readError != nil {
		panic(readError)
	}

	return fileContent
}
