package dictionary

import (
	"encoding/xml"
	"strings"
)

func parseRawDictionary(rawDictionary []byte) rawDictionaryEntries {
	var rawEntries rawDictionaryEntries
	xml.Unmarshal(rawDictionary, &rawEntries)

	return rawEntries
}

func parseDictionary(rawEntries rawDictionaryEntries) []DictionaryEntry {
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

		if entry.PartOfSpeech == nil {
			entries[index].PartOfSpeech = []string{}
		}
	}

	// Some formatting tags are encoded in less-than-ideal way.
	// Search-replace the problematic ones.
	for idx, entry := range entries {
		for defIdx, definition := range entry.Definitions {
			entries[idx].Definitions[defIdx] = strings.Replace(definition, "&quot;", "\"", -1)
		}
	}

	return entries
}

func parseDictionaryEntry(rawEntry rawDictionaryEntry) DictionaryEntry {
	var entry DictionaryEntry

	var partsOfSpeech []string

	for _, feat := range rawEntry.LemmaFormPresentation {
		if strings.EqualFold(feat.Name, "writtenForm") {
			entry.Headword = feat.Value
		}

		if strings.EqualFold(feat.Name, "partOfSpeech") {
			partsOfSpeech = append(partsOfSpeech, feat.Value)
		}

		if strings.EqualFold(feat.Name, "gram") {
			entry.GrammaticalAspect = feat.Value
		}

		if strings.EqualFold(feat.Name, "information") {
			entry.Information = feat.Value
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

	var alternativeForms []string

	for _, wordForm := range rawEntry.WordForm {
		if strings.EqualFold(wordForm.Name, "writtenForm") {
			alternativeForms = append(alternativeForms, wordForm.Value)
		}
	}

	entry.PartOfSpeech = partsOfSpeech
	entry.Definitions = definitions
	entry.AlternativeForms = alternativeForms

	return entry
}
