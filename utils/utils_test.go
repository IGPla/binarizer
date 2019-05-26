package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRemoveTemporaryFiles(t *testing.T) {
	data := []byte("Test")
	filePath := "/tmp/testfile"
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("File should exists. %v", err)
	}

	RemoveTemporaryFiles(&filePath)
	if _, err = os.Stat(filePath); os.IsExist(err) {
		t.Fatalf("File should not exists after deletion. %v", err)
	}
}
