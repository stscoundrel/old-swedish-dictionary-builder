package dictionary

type DictionaryEntry struct {
	Headword          string
	PartOfSpeech      string
	GrammaticalAspect string
	Definitions       []string
	AlternativeForms  []string
}
