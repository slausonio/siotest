package siotest

import (
	"os"
	"path/filepath"
	"testing"
)

// CreateFile test helper that creates a new file with the specified file path.
//
// If an error occurs, it will fail the test.
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

// RemoveFileAndDirs test helper that removes a file and its parent directories.
//
// If an error occurs, it will fail the test.
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
