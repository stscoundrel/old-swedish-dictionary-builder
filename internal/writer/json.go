package writer

import (
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

func WriteMinifiedJson(filePath string, entries []dictionary.DictionaryEntry) {
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
