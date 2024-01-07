package siotest

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/slausonio/go-webserver/environment"
)

// SetEnvVarForTest sets an environment variable for testing purposes.
// It takes a testing.T object, a key string, and a value string as parameters.
// It uses the os.Setenv function to set the environment variable with the provided key and value.
// If there is an error during the setting of the environment variable, it will call the t.Fatal function, passing the error.
// The t.Helper function is called to mark this function as a test helper function and will be skipped during test coverage analysis.
func SetEnvVarForTest(t *testing.T, key, value string) {
	t.Helper()

	err := os.Setenv(key, value)
	if err != nil {
		t.Fatal(err)
	}
}

// SetCurrentEnvForTest sets the current environment to "test" for testing purposes.
// It takes a *testing.T parameter and calls the SetEnvVarForTest function to set the environment variable "environment.CurrentEnv" to "test".
// If an error occurs during setting the environment variable, it will fail the test.
func SetCurrentEnvForTest(t *testing.T) {
	t.Helper()

	SetEnvVarForTest(t, environment.CurrentEnv, "test")
}

// WriteEnvToFile writes the environment variables to a file in the specified filename param.
func WriteEnvToFile(t *testing.T, filename string, env environment.Environment) {
	t.Helper()

	err := godotenv.Write(env, filename)
	if err != nil {
		t.Fatal(err)
	}
}
