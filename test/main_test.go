package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	InitDB("./")

	os.Exit(m.Run())
}
