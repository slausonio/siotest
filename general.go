package siotest

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const CurrentEnvKey = "CURRENT_ENV"

// SetEnvVarForTest test helper that sets an environment variable
//
// It uses the os.Setenv function to set the environment variable with the provided key and value.
//
// If there is an error during the setting of the environment variable, it will fail the test
func SetEnvVarForTest(t *testing.T, key, value string) {
	t.Helper()

	err := os.Setenv(key, value)
	if err != nil {
		t.Fatal(err)
	}
}

// SetCurrentEnvForTest test helper that sets the CURRENT_ENV environment var to "test" for testing purposes.
// calls the SetEnvVarForTest function to set the environment variable "environment.CurrentEnv" to "test".
//
// If an error occurs during setting the environment variable, it will fail the test.
func SetCurrentEnvForTest(t *testing.T) {
	t.Helper()

	SetEnvVarForTest(t, CurrentEnvKey, "test")
}

// WriteEnvToFile test helper that writes the environment variables to a file in the specified filename param.
//
// If an error occurs during setting the environment variable, it will fail the test.
func WriteEnvToFile(t *testing.T, filename string, env map[string]string) {
	t.Helper()

	err := godotenv.Write(env, filename)
	if err != nil {
		t.Fatal(err)
	}
}
