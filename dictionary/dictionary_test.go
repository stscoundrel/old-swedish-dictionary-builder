package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionariesContainExpectedAmountOfEntries(t *testing.T) {
	combined, first, latter := GetOldSwedishDictionary()

	assert.Equal(t, 22572, len(first), "Incorrect amount of entries for first volumes!")
	assert.Equal(t, 19172, len(latter), "Incorrect amount of entries for latter volumes!")
	assert.Equal(t, 41744, len(combined), "Incorrect amount of entries for combined volumes!")
	assert.Equal(t, len(first)+len(latter), len(combined), "Combined should consist of first & latter volumes!")
}

func TestDictionariesContainExpectedContent(t *testing.T) {
	dictionary, _, _ := GetOldSwedishDictionary()

	expected1 := DictionaryEntry{
		Headword:          "af bränna",
		PartOfSpeech:      []string{"vb"},
		GrammaticalAspect: "v.",
		Information:       "",
		Definitions: []string{
			"afbränna, genom eld förstöra. hans trähws the af brendhe  RK 2: 2757 . ib 1511. halff stadhen är affbrändh  BSH 5: 132 (  1506) . Jfr bränna af.",
		},
		AlternativeForms: []string{},
	}

	expected2 := DictionaryEntry{
		Headword:          "alder daghar",
		PartOfSpeech:      []string{"nn"},
		GrammaticalAspect: "pl.",
		Information:       "",
		Definitions: []string{
			"ålderdom.  &quot; tiill aller da[gha] &quot; MD (S) 242 . oppa sina aldher dagha  Lg 3: 650 .",
		},
		AlternativeForms: []string{
			"aller- )",
		},
	}

	expected3 := DictionaryEntry{
		Headword:          "fiädhrdher",
		PartOfSpeech:      []string{"av"},
		GrammaticalAspect: "adj.",
		Information:       "",
		Definitions: []string{
			" Jfr ofiädhradher.",
		},
		AlternativeForms: []string{},
	}

	expected4 := DictionaryEntry{
		Headword:          "ängalund",
		PartOfSpeech:      []string{},
		GrammaticalAspect: "",
		Information:       "",
		Definitions: []string{
			" , se lund.",
		},
		AlternativeForms: []string{},
	}

	assert.Equal(t, expected1, dictionary[100])
	assert.Equal(t, expected2, dictionary[666])
	assert.Equal(t, expected3, dictionary[6666])
	assert.Equal(t, expected4, dictionary[40666])
}
