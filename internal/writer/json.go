package writer

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/stscoundrel/old-swedish-dictionary-builder/dictionary"
)

type MinifiedDictionaryEntry struct {
	Headword          string   `json:"a"`
	PartOfSpeech      []string `json:"b"`
	GrammaticalAspect string   `json:"c"`
	Information       string   `json:"d"`
	Definitions       []string `json:"e"`
	AlternativeForms  []string `json:"f"`
}

func getMinifiedEntries(entries []dictionary.DictionaryEntry) []MinifiedDictionaryEntry {
	// Use alternative struct for minified keys.
	minifiedEntries := []MinifiedDictionaryEntry{}

	for _, entry := range entries {
		minifiedEntry := MinifiedDictionaryEntry{
			entry.Headword,
			entry.PartOfSpeech,
			entry.GrammaticalAspect,
			entry.Information,
			entry.Definitions,
			entry.AlternativeForms,
		}

		minifiedEntries = append(minifiedEntries, minifiedEntry)
	}

	return minifiedEntries
}

func WriteMinifiedJson(filePath string, entries []dictionary.DictionaryEntry) {
	minifiedEntries := getMinifiedEntries(entries)
	file, _ := json.Marshal(minifiedEntries)

	err := ioutil.WriteFile(filePath, file, 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}

func WriteJson(filePath string, entries []dictionary.DictionaryEntry) {
	file, _ := json.MarshalIndent(entries, "", "  ")

	err := ioutil.WriteFile(filePath, file, 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}

func writeGzip(file []byte, fileName string) {
	var fileGZ bytes.Buffer
	zipper := gzip.NewWriter(&fileGZ)

	_, compressErr := zipper.Write(file)
	if compressErr != nil {
		log.Fatalf("Error compressing json file: %+v", compressErr)
	}
	zipper.Close()

	err := ioutil.WriteFile(fileName, fileGZ.Bytes(), 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}

func WriteGzipJson(filePath string, entries []dictionary.DictionaryEntry) {
	file, _ := json.MarshalIndent(entries, "", "  ")
	writeGzip(file, filePath)
}

func WriteMinifiedGzipJson(filePath string, entries []dictionary.DictionaryEntry) {
	minifiedEntries := getMinifiedEntries(entries)
	file, _ := json.MarshalIndent(minifiedEntries, "", "  ")
	writeGzip(file, filePath)
}
