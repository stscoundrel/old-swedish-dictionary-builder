package dictionary

import (
	"encoding/xml"
	"strings"

	"github.com/stscoundrel/old-swedish-dictionary-builder/internal/reader"
)

var VOL_I_TO_III = "fornsvensk"
var VOL_IV_TO__V = "fornsvensk-2"

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

func parseRawDictionary(rawDictionary []byte) RawDictionaryEntries {
	var rawEntries RawDictionaryEntries
	xml.Unmarshal(rawDictionary, &rawEntries)

	return rawEntries
}

func parseDictionary(rawEntries RawDictionaryEntries) []DictionaryEntry {
	var entries []DictionaryEntry

	for _, rawEntry := range rawEntries.Entries {
		entries = append(entries, parseDictionaryEntry(rawEntry))
	}

	return entries
}

func parseDictionaryEntry(rawEntry RawDictionaryEntry) DictionaryEntry {
	var entry DictionaryEntry

	for _, feat := range rawEntry.LemmaFormPresentation {
		if strings.EqualFold(feat.Name, "writtenForm") {
			entry.Headword = feat.Value
		}

		if strings.EqualFold(feat.Name, "partOfSpeech") {
			entry.PartOfSpeech = feat.Value
		}

		if strings.EqualFold(feat.Name, "gram") {
			entry.GrammaticalAspect = feat.Value
		}
	}

	var definitions []string

	for _, sense := range rawEntry.Sense {
		if sense.Prefix.Value != "" {
			definitions = append(definitions, sense.Prefix.Value+" "+sense.Definition.Value)
		} else {
			definitions = append(definitions, sense.Definition.Value)
		}
	}

	entry.Definitions = definitions

	var alternativeForms []string

	for _, wordForm := range rawEntry.WordForm {
		if strings.EqualFold(wordForm.Name, "writtenForm") {
			alternativeForms = append(alternativeForms, wordForm.Value)
		}
	}

	entry.AlternativeForms = alternativeForms

	return entry
}
