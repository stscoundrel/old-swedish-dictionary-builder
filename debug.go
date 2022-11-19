package main

import (
	"fmt"

	"github.com/stscoundrel/old-swedish-dictionary-builder/dictionary"
)

func debug() {
	entries, entries1, entries2 := dictionary.GetOldSwedishDictionary()
	var hasPartOfSpeech int = 0
	var doesNothHavePartOfSpeech int = 0
	var hasGrammaticalAspect int = 0
	var doesNothHaveGrammaticalAspect int = 0

	fmt.Println(len(entries1))
	fmt.Println(len(entries2))
	fmt.Println(len(entries))

	for _, entry := range entries {
		fmt.Println("Headword: " + entry.Headword)
		fmt.Println("GrammaticalAspect: " + entry.GrammaticalAspect)
		fmt.Println("PartOfSpeech: ")
		for _, partOfSpeech := range entry.PartOfSpeech {
			fmt.Println(partOfSpeech)
		}

		fmt.Println("Definitions:")
		for _, definition := range entry.Definitions {
			fmt.Println(definition)
		}

		fmt.Println("Alternative forms:")
		for _, alternativeForm := range entry.AlternativeForms {
			fmt.Println(alternativeForm)
		}

		fmt.Println("########################")

		if len(entry.PartOfSpeech) > 0 {
			hasPartOfSpeech += 1
		} else {
			doesNothHavePartOfSpeech += 1
		}

		if entry.GrammaticalAspect == "" {
			hasGrammaticalAspect += 1
		} else {
			doesNothHaveGrammaticalAspect += 1
		}

	}

	fmt.Println("PART OF SPEECH STATS")
	fmt.Println(hasPartOfSpeech)
	fmt.Println(doesNothHavePartOfSpeech)

	fmt.Println("GRAMMATICAL ASPECT STATS")
	fmt.Println(hasGrammaticalAspect)
	fmt.Println(doesNothHaveGrammaticalAspect)
}
