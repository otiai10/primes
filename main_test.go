package primes

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	teardown(m)
	os.Exit(code)
}

func teardown(m *testing.M) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	if err := os.RemoveAll(filepath.Join(u.HomeDir, ".primes-test")); err != nil {
		panic(err)
	}
}
