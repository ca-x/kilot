package ioplus

import (
	"os"
	"testing"
)

func TestGetWorkingDir(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dir)
}
