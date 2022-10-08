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

type RawDictionaryEntry struct {
	LemmaFormPresentation []Feat  `xml:"Lemma>FormRepresentation>feat"`
	WordForm              []Feat  `xml:"WordForm>feat"`
	Sense                 []Sense `xml:"Sense"`
}

type RawDictionaryEntries struct {
	XMLName xml.Name             `xml:"LexicalResource"`
	Entries []RawDictionaryEntry `xml:"Lexicon>LexicalEntry"`
}
