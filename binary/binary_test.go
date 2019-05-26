package binary

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTestFile(t *testing.T) string {
	data := []byte("echo test")
	filePath := "/tmp/testfile.sh"
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	defer os.Remove(filePath)
	scriptCall := "bash -c"
	newFile := SourceMaker(&filePath, &scriptCall)

	return *newFile
}

func TestSourceMaker(t *testing.T) {
	newFile := createTestFile(t)
	defer os.Remove(newFile)
	if _, err := os.Stat(newFile); os.IsNotExist(err) {
		t.Fatalf("New go source file should exists. %v", err)
	}
}

func TestCompileSource(t *testing.T) {
	newFile := createTestFile(t)
	defer os.Remove(newFile)
	binaryFile := "/tmp/mybinary"
	goos := "linux"
	goarch := "amd64"
	CompileSource(&newFile, &binaryFile, &goos, &goarch)
	defer os.Remove(binaryFile)
	if _, err := os.Stat(binaryFile); os.IsNotExist(err) {
		t.Fatalf("New go binary file should exists. %v", err)
	}
}
