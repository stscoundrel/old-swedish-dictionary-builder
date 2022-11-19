package dictionary

type DictionaryEntry struct {
	Headword          string   `json:"headword"`
	PartOfSpeech      []string `json:"partOfSpeech"`
	GrammaticalAspect string   `json:"grammaticalAspect"`
	Information       string   `json:"information"`
	Definitions       []string `json:"definitions"`
	AlternativeForms  []string `json:"alternativeForms"`
}
