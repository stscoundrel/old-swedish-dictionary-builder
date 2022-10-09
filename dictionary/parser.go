package dictionary

import (
	"encoding/xml"
	"strings"
)

var searchReplaces = map[string]string{
	"\u0026quot;": "'",
	" \u0026":     " &",
}

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
	}

	return entries
}

func parseDictionaryEntry(rawEntry rawDictionaryEntry) DictionaryEntry {
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
		definition := ""
		if sense.Prefix.Value != "" && sense.Definition.Value != "" {
			definition = sense.Prefix.Value + " " + sense.Definition.Value
		} else if sense.Definition.Value != "" {
			definition = sense.Definition.Value
		}

		// Definitions contain encoded special char. Map them to normal outcomes.
		for search, replace := range searchReplaces {
			definition = strings.Replace(definition, search, replace, -1)
		}

		definitions = append(definitions, definition)
	}

	entry.Definitions = definitions

	var alternativeForms []string

	for _, wordForm := range rawEntry.WordForm {
		if strings.EqualFold(wordForm.Name, "writtenForm") {
			alternativeForm := wordForm.Value

			for search, replace := range searchReplaces {
				alternativeForm = strings.Replace(alternativeForm, search, replace, -1)
			}

			alternativeForms = append(alternativeForms, alternativeForm)

		}
	}

	entry.AlternativeForms = alternativeForms

	return entry
}
