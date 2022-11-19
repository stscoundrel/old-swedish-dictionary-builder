package dictionary

type DictionaryEntry struct {
	Headword          string   `json:"headword"`
	PartOfSpeech      []string `json:"partOfSpeech"`
	GrammaticalAspect string   `json:"grammaticalAspect"`
	Definitions       []string `json:"definitions"`
	AlternativeForms  []string `json:"alternativeForms"`
}
