package main

import (
	"errors"
	"os"
	"testing"
)

func TestCreatesJsonBuild(t *testing.T) {
	ToJson()

	// JSON
	_, err1 := os.Stat("./build/old-swedish-dictionary.json")
	_, err2 := os.Stat("./build/old-swedish-dictionary-vol-i-to-iii.json")
	_, err3 := os.Stat("./build/old-swedish-dictionary-vol-iv-to-v.json")
	_, err4 := os.Stat("./build/old-swedish-medieval-law-dictionary.json")

	if errors.Is(err1, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for combined dictionary")
	}

	if errors.Is(err2, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Volumes I to III")
	}

	if errors.Is(err3, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Volumes IV to V")
	}

	if errors.Is(err4, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Schlyters Medieval Law dictionary")
	}
}

func TestCreatesMinifiedJsonBuild(t *testing.T) {
	ToMinifiedJson()

	// Minified JSON
	_, err1 := os.Stat("./build/old-swedish-dictionary.min.json")
	_, err2 := os.Stat("./build/old-swedish-dictionary-vol-i-to-iii.min.json")
	_, err3 := os.Stat("./build/old-swedish-dictionary-vol-iv-to-v.min.json")
	_, err4 := os.Stat("./build/old-swedish-medieval-law-dictionary.min.json")

	if errors.Is(err1, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for combined dictionary")
	}

	if errors.Is(err2, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Volumes I to III")
	}

	if errors.Is(err3, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Volumes IV to V")
	}

	if errors.Is(err4, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Schlyters Medieval Law dictionary")
	}
}

func TestCreatesGzippedJsonBuild(t *testing.T) {
	ToCompressedJson()

	// Compressed JSON
	_, err1 := os.Stat("./build/old-swedish-dictionary.json.gz")
	_, err2 := os.Stat("./build/old-swedish-dictionary-vol-i-to-iii.json.gz")
	_, err3 := os.Stat("./build/old-swedish-dictionary-vol-iv-to-v.json.gz")
	_, err4 := os.Stat("./build/old-swedish-medieval-law-dictionary.json.gz")

	if errors.Is(err1, os.ErrNotExist) {
		t.Error("GZIP JSON output file not found in build directory for combined dictionary")
	}

	if errors.Is(err2, os.ErrNotExist) {
		t.Error("GZIP JSON output file not found in build directory for Volumes I to III")
	}

	if errors.Is(err3, os.ErrNotExist) {
		t.Error("GZIP JSON output file not found in build directory for Volumes IV to V")
	}

	if errors.Is(err4, os.ErrNotExist) {
		t.Error("GZIP JSON output file not found in build directory for Schlyters Medieval Law dictionary")
	}
}

func TestCreatesDSLBuild(t *testing.T) {
	ToDsl()

	// DSL
	_, err1 := os.Stat("./build/old-swedish-dictionary.dsl")
	_, err2 := os.Stat("./build/old-swedish-dictionary-vol-i-to-iii.dsl")
	_, err3 := os.Stat("./build/old-swedish-dictionary-vol-iv-to-v.dsl")
	_, err4 := os.Stat("./build/old-swedish-medieval-law-dictionary.dsl")

	if errors.Is(err1, os.ErrNotExist) {
		t.Error("DSL output file not found in build directory for combined dictionary")
	}

	if errors.Is(err2, os.ErrNotExist) {
		t.Error("DSL output file not found in build directory for Volumes I to III")
	}

	if errors.Is(err3, os.ErrNotExist) {
		t.Error("DSL output file not found in build directory for Volumes IV to V")
	}

	if errors.Is(err4, os.ErrNotExist) {
		t.Error("DSL output file not found in build directory for Schlyters Medieval Law dictionary")
	}
}
