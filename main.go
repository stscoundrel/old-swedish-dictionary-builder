package main

import (
	"github.com/stscoundrel/old-swedish-dictionary-builder/dictionary"
	"github.com/stscoundrel/old-swedish-dictionary-builder/internal/writer"
)

func ToJson() {
	combined, first, latter := dictionary.GetOldSwedishDictionary()

	writer.WriteJson("build/old-swedish-dictionary.json", combined)
	writer.WriteJson("build/old-swedish-dictionary-vol-i-to-iii.json", first)
	writer.WriteJson("build/old-swedish-dictionary-vol-iv-to-v.json", latter)
}

func main() {
	ToJson()
}
