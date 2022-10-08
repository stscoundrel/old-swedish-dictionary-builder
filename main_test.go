package main

import (
	"errors"
	"os"
	"testing"
)

func TestCreatesJsonBuild(t *testing.T) {
	ToJson()

	_, err1 := os.Stat("./build/old-swedish-dictionary.json")
	_, err2 := os.Stat("./build/old-swedish-dictionary-vol-i-to-iii.json")
	_, err3 := os.Stat("./build/old-swedish-dictionary-vol-iv-to-v.json")

	if errors.Is(err1, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for combined dictionary")
	}

	if errors.Is(err2, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Volumes I to III")
	}

	if errors.Is(err3, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Volumes IV to V")
	}
}
