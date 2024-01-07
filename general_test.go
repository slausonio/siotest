package siotest

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSetEnvVarForTest(t *testing.T) {
	t.Helper()
	tests := []struct {
		desc  string
		key   string
		value string
	}{
		{"Regular variable", "VAR_NAME", "VAR_VALUE"},
		{"Empty variable", "EMPTY_NAME", ""},
		{"Numerical variable", "NUM_NAME", "123"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			SetEnvVarForTest(t, tt.key, tt.value)
			got := os.Getenv(tt.key)
			if got != tt.value {
				t.Errorf("SetEnvVarForTest() = %v, want %v", got, tt.value)
			}
		})
	}
}

func TestSetCurrentEnvForTest(t *testing.T) {
	t.Helper()
	SetCurrentEnvForTest(t)
	want := "test"
	got := os.Getenv(CurrentEnvKey)
	if got != want {
		t.Errorf("SetCurrentEnvForTest() = %v, want %v", got, want)
	}
}

func TestWriteEnvToFile(t *testing.T) {
	t.Helper()
	tests := []struct {
		desc     string
		filename string
		env      map[string]string
	}{
		{"Regular file", "file.env", map[string]string{"VAR_NAME": "VAR_VALUE"}},
		{"Empty file", "file-empty.env", map[string]string{}},
		{"Numerical values file", "file-num.env", map[string]string{"NUM_NAME": "123"}},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			WriteEnvToFile(t, tt.filename, tt.env)
			gotEnv, err := godotenv.Read(tt.filename)
			if err != nil {
				t.Fatal(err)
			}
			for key, want := range tt.env {
				got, ok := gotEnv[key]
				if !ok || got != want {
					t.Errorf("WriteEnvToFile() for key %v, got = %v, want %v", key, got, want)
				}
			}

			os.Remove(tt.filename)
		})
	}
}
