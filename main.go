package main

import (
	"github.com/stscoundrel/old-swedish-dictionary-builder/dictionary"
	"github.com/stscoundrel/old-swedish-dictionary-builder/internal/writer"
)

func ToJson() {
	combined, first, latter := dictionary.GetOldSwedishDictionary()
	medievalLaw := dictionary.GetDictionaryOfOldSwedishLaw()

	writer.WriteJson("build/old-swedish-dictionary.json", combined)
	writer.WriteJson("build/old-swedish-dictionary-vol-i-to-iii.json", first)
	writer.WriteJson("build/old-swedish-dictionary-vol-iv-to-v.json", latter)
	writer.WriteJson("build/old-swedish-medieval-law-dictionary.json", medievalLaw)
}

func ToDsl() {
	combined, first, latter := dictionary.GetOldSwedishDictionary()
	medievalLaw := dictionary.GetDictionaryOfOldSwedishLaw()

	writer.WriteDsl("build/old-swedish-dictionary.dsl", combined, "Old Swedish Dictionary")
	writer.WriteDsl("build/old-swedish-dictionary-vol-i-to-iii.dsl", first, "Old Swedish Dictionary Vol I - III")
	writer.WriteDsl("build/old-swedish-dictionary-vol-iv-to-v.dsl", latter, "Old Swedish Dictionary Vol IV - V")
	writer.WriteDsl("build/old-swedish-medieval-law-dictionary.dsl", medievalLaw, "Dictionary of Old Swedish Law")
}

func main() {
	ToJson()
	ToDsl()
}
