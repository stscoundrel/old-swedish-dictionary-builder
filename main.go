package main

import (
	"fmt"

	"github.com/stscoundrel/old-swedish-dictionary-builder/dictionary"
)

func main() {
	entries := dictionary.GetVolumesFourToFive()
	var hasPartOfSpeech int = 0
	var doesNothHavePartOfSpeech int = 0
	var hasGrammaticalAspect int = 0
	var doesNothHaveGrammaticalAspect int = 0

	for _, entry := range entries {
		fmt.Println("Headword: " + entry.Headword)
		fmt.Println("PartOfSpeech: " + entry.PartOfSpeech)
		fmt.Println("GrammaticalAspect: " + entry.GrammaticalAspect)

		fmt.Println("Definitions:")
		for _, definition := range entry.Definitions {
			fmt.Println(definition)
		}

		fmt.Println("Alternative forms:")
		for _, alternativeForm := range entry.AlternativeForms {
			fmt.Println(alternativeForm)
		}

		fmt.Println("########################")

		if entry.PartOfSpeech == "" {
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
