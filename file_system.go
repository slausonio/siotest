package siotest

import (
	"os"
	"path/filepath"
	"testing"
)

// CreateFile creates a new file with the specified file path.
// It also accepts a testing.T pointer for the purpose of testing,
// and uses the t.Helper() method to mark the function as a test helper.
// The function returns a pointer to the created file and logs a fatal error
// if any error occurs during file creation.
func CreateFile(t *testing.T, filePath string) *os.File {
	t.Helper()

	dirPath := filepath.Dir(filePath)

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create directory: %s", err)
	}

	err = os.Chmod(dirPath, 0777)
	if err != nil {
		t.Fatalf("Failed to change directory permissions: %s", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}

	err = os.Chmod(filePath, 0777)
	if err != nil {
		t.Fatalf("Failed to change file permissions: %s", err)
	}

	return file
}

// RemoveFileAndDirs removes a file and its parent directories.
// It also accepts a testing.T pointer for the purpose of testing,
// and uses the t.Helper() method to mark the function as a test helper.
func RemoveFileAndDirs(t *testing.T, filePath string) {
	t.Helper()

	err := os.Remove(filePath)
	if err != nil {
		t.Fatal(err)
	}

	dir := filepath.Dir(filePath)
	err = os.RemoveAll(dir)
	if err != nil {
		t.Fatal(err)
	}
}
