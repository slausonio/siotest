package siotest

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/slausonio/go-webserver/environment"
)

func SetEnvVarForTest(t *testing.T, key, value string) {
	t.Helper()

	err := os.Setenv(key, value)
	if err != nil {
		t.Fatal(err)
	}
}

func SetCurrentEnvForTest(t *testing.T) {
	t.Helper()

	SetEnvVarForTest(t, environment.CurrentEnv, "test")
}

func WriteEnvToFile(t *testing.T, env environment.Environment) {
	t.Helper()

	err := godotenv.Write(env, "env/.env")
	if err != nil {
		t.Fatal(err)
	}
}
