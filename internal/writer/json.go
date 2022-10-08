package writer

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/stscoundrel/old-swedish-dictionary-builder/dictionary"
)

func WriteJson(filePath string, entries []dictionary.DictionaryEntry) {
	file, _ := json.MarshalIndent(entries, "", "  ")

	err := ioutil.WriteFile(filePath, file, 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}
