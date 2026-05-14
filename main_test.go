package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("port", "8084")

	v := getEnv("port", "8080")
	if v != "8084" {
		t.Fatal(fmt.Sprintf("expected: %s, received: %s", "8084", v))
	}

	v = getEnv("host", "localhost")
	if v != "localhost" {
		t.Fatal(fmt.Sprintf("expected: %s, received: %s", "localhost", v))
	}

}
