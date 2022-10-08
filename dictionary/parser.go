package dictionary

import (
	"encoding/xml"
	"strings"
)

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

	// For writing purposes, convert nil slices to empty slices.
	for index, entry := range entries {
		if entry.Definitions == nil {
			entries[index].Definitions = []string{}
		}

		if entry.AlternativeForms == nil {
			entries[index].AlternativeForms = []string{}
		}
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
		if sense.Prefix.Value != "" && sense.Definition.Value != "" {
			definitions = append(definitions, sense.Prefix.Value+" "+sense.Definition.Value)
		} else if sense.Definition.Value != "" {
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
