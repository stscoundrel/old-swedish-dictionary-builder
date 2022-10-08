package dictionary

import (
	"sort"

	"github.com/stscoundrel/old-swedish-dictionary-builder/internal/reader"
)

var VOL_I_TO_III = "fornsvensk"
var VOL_IV_TO__V = "fornsvensk-2"

func combineDictionaries(first []DictionaryEntry, second []DictionaryEntry) ([]DictionaryEntry, []DictionaryEntry, []DictionaryEntry) {
	combined := append(first, second...)

	sort.Slice(combined, func(i, j int) bool {
		return combined[i].Headword < combined[j].Headword
	})

	return combined, first, second
}

func GetVolumesOneToThree() []DictionaryEntry {
	bytes := reader.ReadXmlDictionary(VOL_I_TO_III)

	rawEntries := parseRawDictionary(bytes)
	entries := parseDictionary(rawEntries)

	return entries
}

func GetVolumesFourToFive() []DictionaryEntry {
	bytes := reader.ReadXmlDictionary(VOL_IV_TO__V)

	rawEntries := parseRawDictionary(bytes)
	entries := parseDictionary(rawEntries)

	return entries
}

func GetOldSwedishDictionary() ([]DictionaryEntry, []DictionaryEntry, []DictionaryEntry) {
	firstVolumes := GetVolumesOneToThree()
	latterVolumes := GetVolumesFourToFive()

	return combineDictionaries(firstVolumes, latterVolumes)
}
