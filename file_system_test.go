package siotest

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		hasError bool
	}{
		{
			name:     "ValidFilePath",
			filePath: "tmp/testfile.txt",
			hasError: false,
		},
		{
			name:     "NestedPath",
			filePath: "tmp/subdir/testfile.txt",
			hasError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer RemoveFileAndDirs(t, tt.filePath)

			_ = CreateFile(t, tt.filePath)

			if !tt.hasError {
				if _, err := os.Stat(tt.filePath); os.IsNotExist(err) {
					t.Fatalf("CreateFile() failed to create file at %s", tt.filePath)
				}
			}
		})
	}
}
