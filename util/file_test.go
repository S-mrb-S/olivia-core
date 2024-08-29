package util

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	bytes := FetchFileContent("res/test/test.txt")

	if string(bytes) != "test" {
		t.Errorf("file.ReadFile() failed.")
	}
}
