package dictionary

import "encoding/xml"

type Feat struct {
	Name  string `xml:"att,attr"`
	Value string `xml:"val,attr"`
}

type Sense struct {
	Prefix     Feat `xml:"feat"`
	Definition Feat `xml:"Definition>feat"`
}

type rawDictionaryEntry struct {
	LemmaFormPresentation []Feat  `xml:"Lemma>FormRepresentation>feat"`
	WordForm              []Feat  `xml:"WordForm>feat"`
	Sense                 []Sense `xml:"Sense"`
}

type rawDictionaryEntries struct {
	XMLName xml.Name             `xml:"LexicalResource"`
	Entries []rawDictionaryEntry `xml:"Lexicon>LexicalEntry"`
}
