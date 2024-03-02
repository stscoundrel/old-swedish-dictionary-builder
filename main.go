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

func ToMinifiedJson() {
	combined, first, latter := dictionary.GetOldSwedishDictionary()
	medievalLaw := dictionary.GetDictionaryOfOldSwedishLaw()

	writer.WriteMinifiedJson("build/old-swedish-dictionary.min.json", combined)
	writer.WriteMinifiedJson("build/old-swedish-dictionary-vol-i-to-iii.min.json", first)
	writer.WriteMinifiedJson("build/old-swedish-dictionary-vol-iv-to-v.min.json", latter)
	writer.WriteMinifiedJson("build/old-swedish-medieval-law-dictionary.min.json", medievalLaw)
}

func ToCompressedJson() {
	combined, first, latter := dictionary.GetOldSwedishDictionary()
	medievalLaw := dictionary.GetDictionaryOfOldSwedishLaw()

	writer.WriteGzipJson("build/old-swedish-dictionary.json.gz", combined)
	writer.WriteGzipJson("build/old-swedish-dictionary-vol-i-to-iii.json.gz", first)
	writer.WriteGzipJson("build/old-swedish-dictionary-vol-iv-to-v.json.gz", latter)
	writer.WriteGzipJson("build/old-swedish-medieval-law-dictionary.json.gz", medievalLaw)
}

func ToCompressedMinifiedJson() {
	combined, first, latter := dictionary.GetOldSwedishDictionary()
	medievalLaw := dictionary.GetDictionaryOfOldSwedishLaw()

	writer.WriteMinifiedGzipJson("build/old-swedish-dictionary.min.json.gz", combined)
	writer.WriteMinifiedGzipJson("build/old-swedish-dictionary-vol-i-to-iii.min.json.gz", first)
	writer.WriteMinifiedGzipJson("build/old-swedish-dictionary-vol-iv-to-v.min.json.gz", latter)
	writer.WriteMinifiedGzipJson("build/old-swedish-medieval-law-dictionary.min.json.gz", medievalLaw)

}

func main() {
	ToJson()
	ToDsl()
	ToMinifiedJson()
	ToCompressedJson()
	ToCompressedMinifiedJson()
}
